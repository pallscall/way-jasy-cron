package paladin

import "github.com/go-kratos/kratos/pkg/conf/paladin"

type FilenameGetter interface {
	Filename() string
}

func MustUnmarshalTOML(i interface{}) {
	getter, ok := i.(FilenameGetter)
	if !ok {
		panic(ok)
	}
	filename := getter.Filename()
	if err := paladin.Get(filename).UnmarshalTOML(i); err != nil {
		panic(err)
	}
}
