package tambourine

import (
	"context"
	"fmt"

	"github.com/kris-nova/logger"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/rest"

	"k8s.io/client-go/kubernetes"

	"github.com/kris-nova/novaarchive/filesystem"
	"k8s.io/client-go/tools/clientcmd"
)

// Service will be the main service
// loop for Tambourine
type Service struct {
	logger     Logger
	options    *ServiceOptions
	clientSet  *kubernetes.Clientset
	kubeConfig *rest.Config
}

type ServiceOptions struct {
	KubeConfigPath *filesystem.Path
}

func New(opt *ServiceOptions) *Service {
	return &Service{
		options: opt,
		logger:  &StdOutLogger{},
	}
}

func (s *Service) Connect() error {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", s.options.KubeConfigPath.String())
	if err != nil {
		return fmt.Errorf("unable to load kube config: %v", err)
	}
	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return fmt.Errorf("unable to create kubernetes client: %v", err)
	}
	s.kubeConfig = kubeConfig
	s.clientSet = clientSet
	return nil
}

func (s *Service) GetNodes() (*v1.NodeList, error) {
	return s.clientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
}

func (s *Service) Run() error {
	s.logger.Info("starting service...")
	err := s.Connect()
	if err != nil {
		return fmt.Errorf("unable to Connect(): %v", err)
	}
	nodeList, err := s.GetNodes()
	if err != nil {
		return fmt.Errorf("unable to GetNodes(): %v", err)
	}
	logger.Always("Found nodes: ")
	for _, node := range nodeList.Items {
		logger.Always(node.Name)
	}
	for {
		// TODO Service loop
	}
	return nil
}
