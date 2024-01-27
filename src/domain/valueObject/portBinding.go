package valueObject

import (
	"errors"
	"strings"

	voHelper "github.com/speedianet/control/src/domain/valueObject/helper"
)

type PortBinding struct {
	ServiceName   ServiceName     `json:"serviceName"`
	PublicPort    NetworkPort     `json:"publicPort"`
	ContainerPort NetworkPort     `json:"containerPort"`
	Protocol      NetworkProtocol `json:"protocol"`
	PrivatePort   *NetworkPort    `json:"privatePort"`
}

type serviceBindingInfo struct {
	ServiceNames       []string
	PortBindings       []string
	PublicPortInterval string
}

var httpPublicPortInterval = "80"
var httpsPublicPortInterval = "443"
var databasePublicPortInterval = "30000-39999"

var KnownServiceBindings = []serviceBindingInfo{
	{
		ServiceNames:       []string{"ftp"},
		PortBindings:       []string{"21"},
		PublicPortInterval: "21000-21999",
	},
	{
		ServiceNames:       []string{"ssh", "sftp"},
		PortBindings:       []string{"22"},
		PublicPortInterval: "22000-22999",
	},
	{
		ServiceNames: []string{"telnet"},
		PortBindings: []string{"23"},
	},
	{
		ServiceNames: []string{"dns"},
		PortBindings: []string{"53", "53/udp"},
	},
	{
		ServiceNames: []string{"smtp"},
		PortBindings: []string{"25", "465", "587", "2525"},
	},
	{
		ServiceNames: []string{"whois"},
		PortBindings: []string{"43"},
	},
	{
		ServiceNames: []string{
			"http", "nginx", "caddy", "apache", "httpd", "php",
		},
		PortBindings: []string{"80", "8080"},
	},
	{
		ServiceNames: []string{"kerberos"},
		PortBindings: []string{"88"},
	},
	{
		ServiceNames: []string{"pop3"},
		PortBindings: []string{"110"},
	},
	{
		ServiceNames: []string{"ntp"},
		PortBindings: []string{"123/udp"},
	},
	{
		ServiceNames: []string{"imap"},
		PortBindings: []string{"143"},
	},
	{
		ServiceNames: []string{"ldap"},
		PortBindings: []string{"389"},
	},
	{
		ServiceNames: []string{
			"https",
			"wss",
			"grpcs",
			"php",
			"kong-secure",
		},
		PortBindings: []string{"443", "8443"},
	},
	{
		ServiceNames: []string{"syslog"},
		PortBindings: []string{"514/udp"},
	},
	{
		ServiceNames: []string{"spamassasin", "spam-assassin"},
		PortBindings: []string{"783"},
	},
	{
		ServiceNames: []string{"dot", "dns-over-tls"},
		PortBindings: []string{"853"},
	},
	{
		ServiceNames: []string{"openvpn"},
		PortBindings: []string{"1194"},
	},
	{
		ServiceNames: []string{
			"mssql", "ms-sql", "sqlserver", "sql-server",
		},
		PortBindings:       []string{"1433"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames: []string{
			"oracledb", "oracle-db", "oracle",
		},
		PortBindings:       []string{"1521", "2483", "2484"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames:       []string{"mqtt"},
		PortBindings:       []string{"1883"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames: []string{"nfs"},
		PortBindings: []string{"2049"},
	},
	{
		ServiceNames:       []string{"ghost"},
		PortBindings:       []string{"2368"},
		PublicPortInterval: httpPublicPortInterval,
	},
	{
		ServiceNames:       []string{"node", "nodejs", "ruby-on-rails", "rails", "ruby"},
		PortBindings:       []string{"3000"},
		PublicPortInterval: httpPublicPortInterval,
	},
	{
		ServiceNames:       []string{"mysql", "mariadb", "percona"},
		PortBindings:       []string{"3306"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames:       []string{"nsq"},
		PortBindings:       []string{"4150", "4151", "4160", "4161"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames:       []string{"nats"},
		PortBindings:       []string{"4222", "6222", "8222"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames:       []string{"flask"},
		PortBindings:       []string{"5000"},
		PublicPortInterval: httpPublicPortInterval,
	},
	{
		ServiceNames:       []string{"kibana"},
		PortBindings:       []string{"5601"},
		PublicPortInterval: httpPublicPortInterval,
	},
	{
		ServiceNames:       []string{"postgresql", "postgres"},
		PortBindings:       []string{"5432"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames:       []string{"rabbitmq"},
		PortBindings:       []string{"5672"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames:       []string{"couchdb", "couch"},
		PortBindings:       []string{"5984"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames:       []string{"cassandra"},
		PortBindings:       []string{"9042"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames:       []string{"kafka"},
		PortBindings:       []string{"9092"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames:       []string{"redis"},
		PortBindings:       []string{"6379"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames:       []string{"neo4j"},
		PortBindings:       []string{"7474"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames: []string{
			"django", "kong",
		},
		PortBindings:       []string{"8000"},
		PublicPortInterval: httpPublicPortInterval,
	},
	{
		ServiceNames:       []string{"kong-admin"},
		PortBindings:       []string{"8001"},
		PublicPortInterval: httpPublicPortInterval,
	},
	{
		ServiceNames:       []string{"kong-manager"},
		PortBindings:       []string{"8002"},
		PublicPortInterval: httpPublicPortInterval,
	},
	{
		ServiceNames: []string{"kong-admin-secure"},
		PortBindings: []string{
			"8444",
		},
		PublicPortInterval: httpsPublicPortInterval,
	},
	{
		ServiceNames: []string{"kong-manager-secure"},
		PortBindings: []string{
			"8445",
		},
		PublicPortInterval: httpsPublicPortInterval,
	},
	{
		ServiceNames:       []string{"solr"},
		PortBindings:       []string{"8983"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames: []string{"sonarqube", "sonar"},
		PortBindings: []string{"9000"},
	},
	{
		ServiceNames:       []string{"elasticsearch", "elastic", "elk"},
		PortBindings:       []string{"9200"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames:       []string{"memcached", "memcache"},
		PortBindings:       []string{"11211"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames:       []string{"mongodb", "mongo"},
		PortBindings:       []string{"27017"},
		PublicPortInterval: databasePublicPortInterval,
	},
	{
		ServiceNames: []string{"wireguard"},
		PortBindings: []string{"51820"},
	},
}

func NewPortBinding(
	serviceName ServiceName,
	publicPort NetworkPort,
	containerPort NetworkPort,
	protocol NetworkProtocol,
	privatePort *NetworkPort,
) PortBinding {
	return PortBinding{
		ServiceName:   serviceName,
		PublicPort:    publicPort,
		ContainerPort: containerPort,
		Protocol:      protocol,
		PrivatePort:   privatePort,
	}
}

func findKnownServiceBindingByName(
	receivedServiceName ServiceName,
) (serviceBindingInfo, error) {
	var serviceBinding serviceBindingInfo
	receivedServiceNameStr := receivedServiceName.String()

	desiredServiceBindingIndex := -1

	for bindingIndex, bindingInfo := range KnownServiceBindings {
		standardName := bindingInfo.ServiceNames[0]
		if receivedServiceNameStr == standardName {
			desiredServiceBindingIndex = bindingIndex
			break
		}

		for _, altName := range bindingInfo.ServiceNames {
			if receivedServiceNameStr != altName {
				continue
			}
			desiredServiceBindingIndex = bindingIndex
			break
		}
	}

	if desiredServiceBindingIndex == -1 {
		return serviceBinding, errors.New("UnknownServiceName")
	}

	return KnownServiceBindings[desiredServiceBindingIndex], nil
}

func NewPortBindingsByServiceName(
	receivedServiceName ServiceName,
) ([]PortBinding, error) {
	portBindings := []PortBinding{}

	desiredServiceBinding, err := findKnownServiceBindingByName(receivedServiceName)
	if err != nil {
		return portBindings, err
	}

	for _, portBindingStr := range desiredServiceBinding.PortBindings {
		portBindingParts := strings.Split(portBindingStr, "/")
		if len(portBindingParts) == 0 {
			continue
		}

		publicPort, err := NewNetworkPort(portBindingParts[0])
		if err != nil {
			continue
		}

		containerPort := publicPort

		protocol := GuessNetworkProtocolByPort(publicPort)
		if len(portBindingParts) > 1 && protocol.String() == "tcp" {
			protocol, err = NewNetworkProtocol(portBindingParts[1])
			if err != nil {
				continue
			}
		}

		var privatePortPtr *NetworkPort

		portBinding := NewPortBinding(
			receivedServiceName,
			publicPort,
			containerPort,
			protocol,
			privatePortPtr,
		)
		portBindings = append(portBindings, portBinding)
	}

	return portBindings, nil
}

// format: serviceName[:publicPort][:containerPort][/protocol][:privatePort]
func NewPortBindingFromString(value string) ([]PortBinding, error) {
	portBindings := []PortBinding{}

	value = strings.TrimSpace(value)
	value = strings.ToLower(value)
	portBindingRegex := `^(?P<serviceName>[a-z][\w\.\_\-]{0,128})(?::(?P<publicPort>\d{1,5}))?(?::(?P<containerPort>\d{1,5}))?(?:\/(?P<protocol>\w{1,5}))?(?::(?P<privatePort>\d{1,5}))?$`
	portBindingParts := voHelper.FindNamedGroupsMatches(portBindingRegex, string(value))

	serviceName, err := NewServiceName(portBindingParts["serviceName"])
	if err != nil {
		return portBindings, err
	}

	isServiceUnmapped := false
	servicePortBindings, err := NewPortBindingsByServiceName(serviceName)
	if err != nil {
		isServiceUnmapped = true
	}

	onlyServiceNameSent := portBindingParts["publicPort"] == "" &&
		portBindingParts["containerPort"] == ""
	if onlyServiceNameSent {
		if isServiceUnmapped {
			return portBindings, errors.New("UnknownServiceName")
		}
		return servicePortBindings, nil
	}

	rawPublicPortStr := portBindingParts["publicPort"]
	if rawPublicPortStr == "" {
		return portBindings, errors.New("UnknownPublicPort")
	}

	publicPort, err := NewNetworkPort(rawPublicPortStr)
	if err != nil {
		return portBindings, err
	}

	rawContainerPortStr := portBindingParts["containerPort"]
	if rawContainerPortStr == "" {
		if rawPublicPortStr == "0" {
			return portBindings, errors.New("UnknownContainerPort")
		}
		rawContainerPortStr = rawPublicPortStr
	}

	containerPort, err := NewNetworkPort(rawContainerPortStr)
	if err != nil {
		return portBindings, err
	}

	protocol := GuessNetworkProtocolByPort(publicPort)
	if portBindingParts["protocol"] != "" && protocol.String() == "tcp" {
		protocol, err = NewNetworkProtocol(portBindingParts["protocol"])
		if err != nil {
			return portBindings, err
		}
	}

	var privatePortPtr *NetworkPort
	if portBindingParts["privatePort"] != "" {
		privatePort, err := NewNetworkPort(portBindingParts["privatePort"])
		if err != nil {
			return portBindings, err
		}
		privatePortPtr = &privatePort
	}

	return []PortBinding{
		NewPortBinding(
			serviceName,
			publicPort,
			containerPort,
			protocol,
			privatePortPtr,
		),
	}, nil
}

func (portBinding PortBinding) String() string {
	portBindingStr := portBinding.ServiceName.String()

	if portBinding.PublicPort.String() != "" {
		portBindingStr += ":" + portBinding.PublicPort.String()
	}

	if portBinding.ContainerPort.String() != "" {
		portBindingStr += ":" + portBinding.ContainerPort.String()
	}

	if portBinding.Protocol.String() != "" {
		portBindingStr += "/" + portBinding.Protocol.String()
	}

	if portBinding.PrivatePort != nil {
		portBindingStr += ":" + portBinding.PrivatePort.String()
	}

	return portBindingStr
}

func (portBinding PortBinding) GetPublicPort() NetworkPort {
	return portBinding.PublicPort
}

func (portBinding PortBinding) GetContainerPort() NetworkPort {
	return portBinding.ContainerPort
}

func (portBinding PortBinding) GetProtocol() NetworkProtocol {
	return portBinding.Protocol
}

func (portBinding PortBinding) GetPublicPortInterval() (string, error) {
	serviceInfo, err := findKnownServiceBindingByName(portBinding.ServiceName)
	if err != nil {
		return "", err
	}
	return serviceInfo.PublicPortInterval, nil
}
