package k8s

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseComponentImage(t *testing.T) {
	type args struct {
		kubectlExecOutput string
		imageSection      string
	}
	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{"when getComponentImageTag is passed with a valid output and imageSection and it returns the image tag",
			args{"'my-hash.dkr.ecr.eu-west-1.amazonaws.com/amazon-k8s-cni:my-version'", "imageTag"},
			"my-version", nil},
		{"when getComponentImageTag is passed with a valid output and imagePrefix and it returns the image tag",
			args{"'my-hash.dkr.ecr.eu-west-1.amazonaws.com/amazon-k8s-cni:my-version'", "imagePrefix"},
			"my-hash.dkr.ecr.eu-west-1.amazonaws.com/amazon-k8s-cni", nil},
		{"when getComponentImageTag is passed with a valid output and imagePrefix and it returns the image tag",
			args{"'my-hash.dkr.ecr.eu-west-1.amazonaws.com/amazon-k8s-cni:my-version'", "foo"},
			"", errors.New("invalid imageSection Passed")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseComponentImage(tt.args.kubectlExecOutput, tt.args.imageSection)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, tt.err, err)
		})
	}
}
