package superfile

import "github.com/helloticket/superfile/model"

type LayoutFake struct {
}

func (l *LayoutFake) GetSegmentoDefinitions() model.RecordDetailMap {
	return model.RecordDetailMap{}
}

func (l *LayoutFake) GetRemessaLayout() model.FileConfigMap {
	return model.FileConfigMap{}
}

func (l *LayoutFake) GlobalSettings() map[string]string {
	return map[string]string{}
}

func (l *LayoutFake) GetRetornoLayout() model.FileConfigMap {
	return model.FileConfigMap{}
}

func (l *LayoutFake) GetVersao() string {
	return ""
}

func (l *LayoutFake) GetServico() string {
	return ""
}

func (l *LayoutFake) GetLayout() string {
	return ""
}

func (l *LayoutFake) GetTamanhoRegistro() int64 {
	return 0
}

func (l *LayoutFake) Validate() error {
	return nil
}
