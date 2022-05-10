package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ijwfly/go-particles-experiment/objects/particle"
	"github.com/ijwfly/go-particles-experiment/objects/particlesource"
	"github.com/ijwfly/go-particles-experiment/primitives"
	"log"
	"math/rand"
	"time"
)

const (
	screenWidth  = 1800
	screenHeight = 1024
	particlesNum = 0
)

type Game struct {
	particles       []particle.Particle
	particleSource1 *particlesource.ParticleSource
	particleSource2 *particlesource.ParticleSource
	keys            []ebiten.Key
}

func NewGame() *Game {
	g := &Game{}
	for i := 0; i < particlesNum; i++ {
		position := primitives.NewPoint(rand.Float64()*screenWidth, rand.Float64()*screenHeight)
		velocity := primitives.RandomNormalizedVector().MultiplyBy(5)
		currentParticle := particle.NewParticle(position, velocity)
		currentParticle.Init()
		g.particles = append(g.particles, *currentParticle)
	}

	particleSource1Position := primitives.NewPoint(screenWidth*20/100, screenHeight/2)
	particleSource1Direction := primitives.NewVector(1, 0)
	g.particleSource1 = particlesource.NewParticleSource(particleSource1Position, particleSource1Direction, screenHeight/10)
	g.particleSource1.Init()

	particleSource2Position := primitives.NewPoint(screenWidth*80/100, screenHeight/2)
	particleSource2Direction := primitives.NewVector(-1, 0)
	g.particleSource2 = particlesource.NewParticleSource(particleSource2Position, particleSource2Direction, screenHeight/10)
	g.particleSource2.Init()

	return g
}

func (g *Game) Update() error {

	allParticles := append(g.particles, g.particleSource1.Particles()...)
	allParticles = append(allParticles, g.particleSource2.Particles()...)

	for i := 0; i < len(allParticles); i++ {
		currentParticle := allParticles[i]

		for j := 0; j < len(allParticles); j++ {
			if i == j {
				continue
			}
			otherParticle := allParticles[j]
			// calculate the distance between the two particles
			distance := currentParticle.Position.DistanceTo(otherParticle.Position)
			forceValue := 0.01 / (distance * distance)
			// calculate the force vector
			force := primitives.NewVector(currentParticle.Position.X-otherParticle.Position.X, currentParticle.Position.Y-otherParticle.Position.Y).Normalize().MultiplyBy(forceValue)
			// add the force to the currentParticle
			currentParticle.AddForce(force)
		}

		// add friction
		currentParticle.Velocity.MultiplyBy(0.99)

		//move particle back to screen
		//if currentParticle.Position.X < 0 {
		//	currentParticle.Position.X = screenWidth
		//} else if currentParticle.Position.X > screenWidth {
		//	currentParticle.Position.X = 0
		//}
		//if currentParticle.Position.Y < 0 {
		//	currentParticle.Position.Y = screenHeight
		//} else if currentParticle.Position.Y > screenHeight {
		//	currentParticle.Position.Y = 0
		//}

		currentParticle.Update()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.particleSource1.Start()
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		g.particleSource1.Stop()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.particleSource2.Start()
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyRight) {
		g.particleSource2.Stop()
	}

	g.particleSource1.Update()
	g.particleSource2.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	allParticles := append(g.particles, g.particleSource1.Particles()...)
	allParticles = append(allParticles, g.particleSource2.Particles()...)
	for _, currentParticle := range allParticles {
		currentParticle.Draw(screen)
	}

	g.particleSource1.Draw(screen)
	g.particleSource2.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Particles")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
