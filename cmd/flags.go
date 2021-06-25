package cmd

import (
	"github.com/urfave/cli/v2"
)

// Variables declared for CLI
var (
	AddServer        string
	ViewImages       string
	CreateVM         string
	ContainerName    string
	Ports            string
	Mode             string
	RemoveVM         string
	ID               string
	Specs            string
	GPU              bool
	UpdateServerList bool
	ServerList       bool
	SetDefaultConfig bool
	NetworkInterface bool
)

var AppConfigFlags = []cli.Flag{
	// Deprecated to be implemented using GRPC
	&cli.StringFlag{
		Name:        "Mode",
		Value:       "client",
		Usage:       "Specifies mode of running",
		EnvVars: []string{"P2P_MODE"},
		Destination: &Mode,
	},
	&cli.BoolFlag{
		Name:        "UpdateServerList",
		Usage:       "Update List of Server available based on servers iptables",
		EnvVars: []string{"UPDATE_SERVER_LIST"},
		Destination: &UpdateServerList,
	},
	&cli.BoolFlag{
		Name:        "ListServers",
		Usage:       "List servers which can render tasks",
		EnvVars: []string{"LIST_SERVERS"},
		Destination: &ServerList,
	},
	&cli.StringFlag{
		Name:        "AddServer",
		Usage:       "Adds server IP address to iptables",
		EnvVars: []string{"ADD_SERVER"},
		Destination: &AddServer,
	},
	&cli.StringFlag{
		Name:        "ViewImages",
		Usage:       "View images available on the server IP address",
		EnvVars: []string{"VIEW_IMAGES"},
		Destination: &ViewImages,
	},
	&cli.StringFlag{
		Name:        "CreateVM",
		Usage:       "Creates Docker container on the selected server",
		EnvVars: []string{"CREATE_VM"},
		Destination: &CreateVM,
	},
	&cli.StringFlag{
		Name:        "ContainerName",
		Usage:       "Specifying the container run on the server side",
		EnvVars: []string{"CONTAINER_NAME"},
		Destination: &ContainerName,
	},
	&cli.StringFlag{
		Name:        "RemoveVM",
		Usage:       "Stop and Remove Docker container",
		EnvVars: []string{"REMOVE_VM"},
		Destination: &RemoveVM,
	},
	&cli.StringFlag{
		Name:        "ID",
		Usage:       "Docker Container ID",
		EnvVars: []string{"ID"},
		Destination: &ID,
	},
	&cli.StringFlag{
		Name:        "Ports",
		Usage:       "Number of ports to open for the Docker Container",
		EnvVars: []string{"NUM_PORTS"},
		Destination: &Ports,
	},
	&cli.BoolFlag{
		Name:        "GPU",
		Usage:       "Create Docker Containers to access GPU",
		EnvVars: []string{"USE_GPU"},
		Destination: &GPU,
	},
	&cli.StringFlag{
		Name:        "Specs",
		Usage:       "Specs of the server node",
		EnvVars: []string{"SPECS"},
		Destination: &Specs,
	},

	&cli.BoolFlag{
		Name:        "SetDefaultConfig",
		Usage:       "Sets a default configuration file",
		EnvVars: []string{"SET_DEFAULT_CONFIG"},
		Destination: &SetDefaultConfig,
	},
	&cli.BoolFlag{
		Name:        "NetworkInterfaces",
		Usage:       "Shows the network interface in your computer",
		EnvVars: []string{"NETWORK_INTERFACE"},
		Destination: &NetworkInterface,
	},
}