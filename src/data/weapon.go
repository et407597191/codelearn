package data

// Weapopn 武器类
type Weapopn interface {
	Speed() int          // 移动速度
	AttackDistance() int // 攻击距离
	ViewDistance() int   // 视野距离
	Energy() int         // 能量值
	AttackPower() int    // 攻击力
}

type BaseWeapon struct {
	speed         int
	attckDistance int
	viewDistance  int
	energy        int
	attackPower   int
}

func (b *BaseWeapon) Speed() int {
	return b.speed
}

func (b *BaseWeapon) AttackDistance() int {
	return b.attckDistance
}

func (b *BaseWeapon) Energy() int {
	return b.energy
}

func (b *BaseWeapon) ViewDistance() int {
	return b.viewDistance
}

func (b *BaseWeapon) AttackPower() int {
	return b.attackPower
}
