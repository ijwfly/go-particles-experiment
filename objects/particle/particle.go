package particle

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gogame/m/v2/primitives"
	"image/color"
)

const particleMaxSpeed = 30

type Particle struct {
	Position     *primitives.Point
	TailPosition *primitives.Point
	Velocity     *primitives.Vector
	Acceleration *primitives.Vector
}

func NewParticle(position *primitives.Point, velocity *primitives.Vector) *Particle {
	p := &Particle{
		Position: position,
		Velocity: velocity,
	}
	p.TailPosition = primitives.NewPoint(p.Position.X-p.Velocity.X, p.Position.Y-p.Velocity.Y)
	p.Acceleration = primitives.NewVector(0, 0)
	return p
}

func (p *Particle) AddForce(force *primitives.Vector) {
	p.Acceleration.Add(force)
}

func (p *Particle) Init() {
	p.TailPosition = primitives.NewPoint(p.Position.X-p.Velocity.X, p.Position.Y-p.Velocity.Y)
}

func (p *Particle) Update() {
	p.Velocity.Add(p.Acceleration)
	if p.Velocity.Length() > particleMaxSpeed {
		p.Velocity.Normalize().MultiplyBy(particleMaxSpeed)
	}

	// move the particle
	p.Position.MoveByVector(p.Velocity)

	// move the tail
	p.TailPosition.Update(p.Position.X-p.Velocity.X, p.Position.Y-p.Velocity.Y)
}

func (p *Particle) Draw(screen *ebiten.Image) {
	// draw the particle
	c := color.RGBA{R: uint8(0xbb),
		G: uint8(0xdd),
		B: uint8(0xff),
		A: 0xff}
	ebitenutil.DrawLine(screen, p.Position.X, p.Position.Y, p.TailPosition.X, p.TailPosition.Y, c)
}
