package advancedlog

import (
	"github.com/sirupsen/logrus"
)

type Closer func(*logrus.Entry) error
