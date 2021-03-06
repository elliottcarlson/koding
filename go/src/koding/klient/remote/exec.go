package remote

import (
	"errors"
	"fmt"

	"github.com/koding/kite"

	"koding/klient/remote/req"
)

// ExecHandler runs the given command on the given remote klient.
func (r *Remote) ExecHandler(kreq *kite.Request) (interface{}, error) {
	log := r.log.New("remote.exec")

	var params req.Exec

	if kreq.Args == nil {
		return nil, errors.New("Required arguments were not passed.")
	}

	if err := kreq.Args.One().Unmarshal(&params); err != nil {
		err = fmt.Errorf(
			"remote.sshKeyAdd: Error '%s' while unmarshalling request '%s'\n",
			err, kreq.Args.One(),
		)

		log.Error(err.Error())
		return nil, err
	}

	switch {
	case params.Machine == "":
		return nil, errors.New("Missing required argument `machine`.")
	case params.Command == "":
		return nil, errors.New("Missing required argument `command`.")
	}

	log = log.New(
		"machineName", params.Machine,
	)

	remoteMachine, err := r.GetDialedMachine(params.Machine)
	if err != nil {
		log.Error("Error getting dialed, valid machine. err:%s", err)
		return nil, err
	}

	// cd into path and then run the command if path is not nil.
	var cmd = params.Command
	if params.Path != "" {
		cmd = fmt.Sprintf("cd %s && %s", params.Path, params.Command)
	}

	var execReq = struct{ Command string }{cmd}

	return remoteMachine.Tell("exec", execReq)
}
