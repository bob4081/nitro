package nitro

import (
	"fmt"

	"github.com/craftcms/nitro/config"
	"github.com/craftcms/nitro/validate"
)

// CreateDatabaseContainer is responsible for the creation of a new Docker database and will
// assign a volume and port based on the arguments. Validation of port collisions should occur
// outside of this func and this will only validate engines and versions.
func CreateDatabaseContainer(machine, engine, version, port string) (*Action, error) {
	if err := validate.DatabaseEngineAndVersion(engine, version); err != nil {
		return nil, err
	}

	// get the container path and port based on the engine
	var containerPath string
	var containerPort string
	var containerConfPath string
	var hostConfPath string
	var containerEnvVars []string
	switch engine {
	case "postgres":
		containerPort = "5432"
		containerPath = "/var/lib/postgresql/data"
		hostConfPath = "/home/ubuntu/.nitro/databases/postgres/conf.d/"
		containerConfPath = "/etc/mysql/conf.d"
		containerEnvVars = []string{"-e", "POSTGRES_PASSWORD=nitro", "-e", "POSTGRES_USER=nitro"}
	default:
		containerPort = "3306"
		containerPath = "/var/lib/mysql"
		hostConfPath = "/home/ubuntu/.nitro/databases/mysql/conf.d/"
		containerConfPath = "/etc/mysql/conf.d"
		containerEnvVars = []string{"-e", "MYSQL_ROOT_PASSWORD=nitro", "-e", "MYSQL_USER=nitro", "-e", "MYSQL_PASSWORD=nitro"}
	}

	// create the volumeMount path using the engine, version, and port
	volume := containerVolume(engine, version, port)
	volumeMount := fmt.Sprintf("%s:%s", volume, containerPath)

	// build the container machine based on engine, version, and port
	containerName := containerName(engine, version, port)

	// create the port mapping
	portMapping := fmt.Sprintf("%v:%v", port, containerPort)

	args := []string{"exec", machine, "--", "docker", "run", "-v", hostConfPath + ":" + containerConfPath, "-v", volumeMount, "--name", containerName, "-d", "--restart=always", "-p", portMapping}

	// append the env vars
	args = append(args, containerEnvVars...)

	// append the image and tag
	args = append(args, engine+":"+version)

	return &Action{
		Type:       "exec",
		UseSyscall: false,
		Args:       args,
	}, nil
}

// SetDatabaseUserPermissions is used to set all permissions on the nitro user for a database
func SetDatabaseUserPermissions(machine string, database config.Database) (*Action, error) {
	return &Action{
		Type:       "exec",
		UseSyscall: false,
		Args:       []string{"exec", machine, "--", "sudo", "bash", "/opt/nitro/scripts/docker-set-database-user-permissions.sh", database.Name(), database.Engine, ">", "/dev/null", "2>&1"},
	}, nil
}

// CreateDatabaseVolume will make a database vaolume to ensure that data is persisted during reboots.
func CreateDatabaseVolume(machine, engine, version, port string) (*Action, error) {
	if err := validate.DatabaseEngineAndVersion(engine, version); err != nil {
		return nil, err
	}

	volume := containerVolume(engine, version, port)

	return &Action{
		Type:       "exec",
		UseSyscall: false,
		Args:       []string{"exec", machine, "--", "docker", "volume", "create", volume},
	}, nil
}

func containerName(engine, version, port string) string {
	return fmt.Sprintf("%s_%s_%s", engine, version, port)
}

func containerVolume(engine, version, port string) string {
	return fmt.Sprintf("%s_%s_%s", engine, version, port)
}
