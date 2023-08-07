package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	jsonserializer "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/kubectl/pkg/scheme"
)

var (
	ETCDCaFile  string = "/etc/kubernetes/pki/etcd/ca.crt"
	ETCDCrtFile string = "/etc/kubernetes/pki/etcd/server.crt"
	ETCDKeyFile string = "/etc/kubernetes/pki/etcd/server.key"
)

type EtcdCertificateConfig struct {
	etcdtls      tls.Certificate
	etcdCertPool *x509.CertPool
}

type etcd3kv struct {
	Key            string `json:"key,omitempty"`
	Value          string `json:"value,omitempty"`
	CreateRevision int64  `json:"create_revision,omitempty"`
	ModRevision    int64  `json:"mod_revision,omitempty"`
	Version        int64  `json:"version,omitempty"`
	Lease          int64  `json:"lease,omitempty"`
}

func loadEtcdCertificate(ecc *EtcdCertificateConfig) {
	// 加载CA证书、客户端证书和私钥文件
	caCert, err := os.ReadFile(ETCDCaFile)
	if err != nil {
		log.Fatal(err)
	}
	clientCert, err := os.ReadFile(ETCDCrtFile)
	if err != nil {
		log.Fatal(err)
	}
	clientKey, err := os.ReadFile(ETCDKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	// 创建TLS配置
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	cert, err := tls.X509KeyPair(clientCert, clientKey)
	if err != nil {
		log.Fatal(err)
	}

	ecc.etcdCertPool = certPool
	ecc.etcdtls = cert
}

func formatPrintJson(jsonData []byte) {
	// 创建一个空的 interface{} 变量来存储解析后的 JSON 数据
	var resultJson interface{}

	// 解析 JSON 数据到 interface{}
	err := json.Unmarshal(jsonData, &resultJson)
	if err != nil {
		fmt.Println("解析 JSON 失败：", err)
		return
	}

	// 将格式化后的 JSON 数据打印输出
	formattedJSON, err := json.MarshalIndent(resultJson, "", "  ")
	if err != nil {
		fmt.Println("格式化 JSON 失败：", err)
		return
	}
	fmt.Printf("%s \n", string(formattedJSON))

	// TODO ...
}

func etcdDump(client *clientv3.Client) error {
	response, err := clientv3.NewKV(client).Get(context.Background(), "/registry/pods/tfsk/envoy-f4dcfcf47-bcxns", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		return err
	}

	kvData := []etcd3kv{}
	decoder := scheme.Codecs.UniversalDeserializer()
	encoder := jsonserializer.NewSerializer(jsonserializer.DefaultMetaFactory, scheme.Scheme, scheme.Scheme, false)
	objJSON := &bytes.Buffer{}

	for _, kv := range response.Kvs {
		obj, _, err := decoder.Decode(kv.Value, nil, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "WARN: error decoding value %q: %v\n", string(kv.Value), err)
			continue
		}
		objJSON.Reset()
		if err := encoder.Encode(obj, objJSON); err != nil {
			fmt.Fprintf(os.Stderr, "WARN: error encoding object %#v as JSON: %v", obj, err)
			continue
		}
		kvData = append(
			kvData,
			etcd3kv{
				Key:            string(kv.Key),
				Value:          string(objJSON.Bytes()),
				CreateRevision: kv.CreateRevision,
				ModRevision:    kv.ModRevision,
				Version:        kv.Version,
				Lease:          kv.Lease,
			},
		)
	}

	jsonData, err := json.MarshalIndent(kvData, "", "  ")
	if err != nil {
		return err
	}

	formatPrintJson(jsonData)
	//fmt.Println(string(jsonData))

	return nil
}

func main() {
	// 加载etcd服务端ca证书 以及 秘钥对
	var etcdCertConfig EtcdCertificateConfig
	loadEtcdCertificate(&etcdCertConfig)

	tlsConfig := &tls.Config{
		RootCAs:      etcdCertConfig.etcdCertPool,
		Certificates: []tls.Certificate{etcdCertConfig.etcdtls},
	}

	// 创建etcd客户端连接
	etcdcli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"https://127.0.0.1:2379"}, // 替换为您的etcd地址
		TLS:         tlsConfig,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer etcdcli.Close()

	err = etcdDump(etcdcli)
	if err != nil {
		log.Fatal(err)
	}
}
