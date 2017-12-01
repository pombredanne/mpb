package mpb

import "github.com/vbauerster/mpb/decor"

// BarOption is a function option which changes the default behavior of a bar,
// if passed to p.AddBar(int64, ...BarOption)
type BarOption func(*bState)

// AppendDecorators let you inject decorators to the bar's right side
func AppendDecorators(appenders ...decor.DecoratorFunc) BarOption {
	return func(s *bState) {
		s.appendFuncs = append(s.appendFuncs, appenders...)
	}
}

// PrependDecorators let you inject decorators to the bar's left side
func PrependDecorators(prependers ...decor.DecoratorFunc) BarOption {
	return func(s *bState) {
		s.prependFuncs = append(s.prependFuncs, prependers...)
	}
}

// BarTrimLeft trims left side space of the bar
func BarTrimLeft() BarOption {
	return func(s *bState) {
		s.trimLeftSpace = true
	}
}

// BarTrimRight trims right space of the bar
func BarTrimRight() BarOption {
	return func(s *bState) {
		s.trimRightSpace = true
	}
}

// BarTirm trims both left and right spaces of the bar
func BarTrim() BarOption {
	return func(s *bState) {
		s.trimLeftSpace = true
		s.trimRightSpace = true
	}
}

// BarID overwrites internal bar id
func BarID(id int) BarOption {
	return func(s *bState) {
		s.id = id
	}
}

// BarEtaAlpha option is a way to adjust ETA behavior.
// You can play with it, if you're not satisfied with default behavior.
// Default value is 0.25.
func BarEtaAlpha(a float64) BarOption {
	return func(s *bState) {
		s.etaAlpha = a
	}
}

// BarDropRatio sets drop ratio, default is 10. Effective when total is dynamic.
// If progress tip reaches total, but total is not final value yet, tip will be
// dropped by specified ratio.
func BarDropRatio(ratio int64) BarOption {
	return func(s *bState) {
		s.dropRatio = ratio
	}
}

func barWidth(w int) BarOption {
	return func(s *bState) {
		s.width = w
	}
}

func barFormat(format string) BarOption {
	return func(s *bState) {
		s.updateFormat(format)
	}
}
