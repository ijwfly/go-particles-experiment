package particlesource

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gogame/m/v2/objects/particle"
	"gogame/m/v2/primitives"
	"image/color"
	"math/rand"
)

const particleLifetime = 120
const particleEmissionRate = 3

type emittedParticle struct {
	particle *particle.Particle
	lifeTime int
}

type ParticleSource struct {
	Position  *primitives.Point
	Direction *primitives.Vector
	Radius    float64

	running   bool
	particles []emittedParticle
}

func NewParticleSource(position *primitives.Point, direction *primitives.Vector, radius float64) *ParticleSource {
	return &ParticleSource{
		Position:  position,
		Direction: direction,
		Radius:    radius,
	}
}

func (ps *ParticleSource) Init() {
	ps.Direction.Normalize()
	ps.particles = make([]emittedParticle, 0, particleEmissionRate*particleLifetime)
}

func (ps *ParticleSource) Start() {
	ps.running = true
}

func (ps *ParticleSource) Stop() {
	ps.running = false
}

func (ps *ParticleSource) Update() {
	for i := 0; i < len(ps.particles); i++ {
		ps.particles[i].lifeTime--
		if ps.particles[i].lifeTime <= 0 {
			ps.particles = append(ps.particles[:i], ps.particles[i+1:]...)
		}
	}

	if ps.running {
		newParticles := make([]emittedParticle, 0, particleEmissionRate)
		for i := 0; i < particleEmissionRate; i++ {
			particlePosition := *ps.Position
			particlePosition.MoveByVector(primitives.RandomNormalizedVector().MultiplyBy(rand.Float64() * ps.Radius))
			newParticle := particle.NewParticle(&particlePosition, primitives.MultiplyVectorBy(ps.Direction, rand.Float64()*12))
			newParticles = append(newParticles, emittedParticle{newParticle, particleLifetime})
		}
		ps.particles = append(ps.particles, newParticles...)
	}
}

func (ps *ParticleSource) Draw(screen *ebiten.Image) {
	// draw the source
	c := color.RGBA{R: uint8(0xbb),
		G: uint8(0xdd),
		B: uint8(0xff),
		A: 0xff}
	lineto := *ps.Position
	lineto.MoveByVector(primitives.MultiplyVectorBy(ps.Direction, ps.Radius))
	ebitenutil.DrawLine(screen, ps.Position.X, ps.Position.Y, lineto.X, lineto.Y, c)
}

func (ps *ParticleSource) Particles() []particle.Particle {
	particles := make([]particle.Particle, 0, len(ps.particles))
	for _, p := range ps.particles {
		particles = append(particles, *p.particle)
	}
	return particles
}
