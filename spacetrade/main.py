import pygame
import title
import planets
import planet
import random

## uses a tiny amount of boilerplate code
## this is a terrible rewrite/port of my original js code for my go engine
## the original game broke and was not worth my time


pygame.init()
screen_size = (640, 480)

screen = pygame.display.set_mode(screen_size)
pygame.display.set_caption("Commercium")

clock = pygame.time.Clock()

font = pygame.font.Font('freesansbold.ttf', 32) 

running = True

ts = title.titleScreen()

planetsData = {
	"Earth": ["H2O", "Si", "U"],
	"Mars": ["Fe", "Ti", "Eu"],
	"Venus": ["C", "N", "O"]
}

itemsData = {
	"H2O": 130.0,
	"Fe": 20.0,
	"Si": 60.0,
	"U": 200.0,
	"Eu": 500.0,
	"Ti": 50.0,
	"C": 30.0,
	"N": 40.0,
	"O": 110.0
}

class inventoryManager:
	def __init__(self):
		self.bankBalance = 100.0
		self.inventoryData = {
			"H2O": 0.0,
			"Fe":  0.0,
			"Si": 0.0,
			"U": 0.0,
			"Eu": 0.0,
			"Ti": 0.0,
			"C": 0.0,
			"N": 0.0,
			"O": 0.0,
			"I AM RICH": 0.0
		}
	def incItem(self, name):
		self.inventoryData[name] = self.inventoryData[name] + 1
	def decItem(self, name):
		self.inventoryData[name] = self.inventoryData[name] - 1
	def setBalance(self, amount):
		self.bankBalance = amount

class inGameService:
	def __init__(self):
		self.isInGame = False
		self.isInPlanet = False
		self.isInSpaceStation = False
		self.planetName = ""
		self.counterMax = 1800
	def setInGame(self, b):
		self.isInGame = b
	def setInPlanet(self, b):
		self.isInPlanet = b
	def setInSpaceStation(self, b):
		self.isInSpaceStation = b
	def setPlanetName(self, name):
		self.planetName = name
	def setCounterMax(self, i):
		self.counterMax = i

inGame = inGameService()
invMan = inventoryManager()

counterForNextPriceUpdate = 0

bgImage = pygame.image.load('Planet-1.png')

while running:
	for event in pygame.event.get():
		if event.type == pygame.QUIT:
			running = False
		elif event.type == pygame.KEYDOWN:
			if event.key == pygame.K_ESCAPE:
				if inGame.isInPlanet:
					inGame.setInPlanet(False)
				elif inGame.isInSpaceStation:
					inGame.setInSpaceStation(False)
	screen.fill((0,0,0))
	#background
	mousepos = pygame.mouse.get_pos()
	xoffset = ((mousepos[0]-160)/640)*10
	yoffset = ((mousepos[1]-120)/480)*10
	screen.blit(bgImage, (0+xoffset,0+yoffset))
	if inGame.isInGame:
		counterForNextPriceUpdate = counterForNextPriceUpdate + 1
		#update prices every x period of time
		if counterForNextPriceUpdate == inGame.counterMax:
			#prices
			for itemName in itemsData:
				while True:
					#modify the item price by a random number (neg or pos)
					proposedPrice = itemsData[itemName] + random.randint(-50,50)
					#out of bounds check
					if proposedPrice > 0:
						itemsData[itemName] = proposedPrice
						break
			#reset counter
			counterForNextPriceUpdate = 0
		#drawing/interactivity
		if not inGame.isInPlanet and not inGame.isInSpaceStation:
			planets.drawPlanets(screen, inGame, invMan)
		if inGame.isInPlanet:
			planet.drawPlanet(screen, inGame, planetsData, itemsData, invMan)
		if inGame.isInSpaceStation:
			planet.drawSpaceStation(screen, inGame, invMan)
	else:
		ts.drawTitleScreen(screen, font, inGame)

	pygame.display.flip()
	clock.tick(60)

pygame.quit()