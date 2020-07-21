import pygame
import gui

earthImage = pygame.image.load('earth.png')
marsImage = pygame.image.load('mars.png')
venusImage = pygame.image.load('venus.png')

def drawPlanets(screen, inGame, invMan):
	#title
	font = pygame.font.Font('freesansbold.ttf', 25) 
	text = font.render("IMT - Interplanetary Market Terminal", True, (255,255,255)) 
	textRect = text.get_rect()  
	textRect.topleft = (20, 20)
	screen.blit(text, textRect)
	#slot 1
	drawPlanetCardWithImage(screen, "Earth", earthImage, (100, 200), inGame)
	#slot 2
	drawPlanetCardWithImage(screen, "Mars", marsImage, (280, 200), inGame)
	#slot 3
	drawPlanetCardWithImage(screen, "Venus", venusImage, (460, 200), inGame)
	#balance
	gui.button(screen, "$"+str(invMan.bankBalance), (640 - 80, 30))
	#space station
	if(gui.button(screen, "Space Station", (120, 440))):
		inGame.setInSpaceStation(True)

def drawPlanetCard(screen, name, pos, inGame):
	bgColour = (255,255,255)
	#check for hover state
	mousepos = pygame.mouse.get_pos()
	if pygame.Rect(pos, (80, 80)).collidepoint(mousepos):
		bgColour = (84,84,255)
		#detect mouse click
		if pygame.mouse.get_pressed()[0]:
			inGame.setInPlanet(True)
			inGame.setPlanetName(name)
	font = pygame.font.Font('freesansbold.ttf', 25) 
	text = font.render(name, True, bgColour) 
	textRect = text.get_rect()  
	textRect.center = (pos[0] + 40, pos[1] - 40) 
	screen.blit(text, textRect)
	pygame.draw.rect(screen, bgColour, (pos[0], pos[1], 80, 80))

def drawPlanetCardWithImage(screen, name, image, pos, inGame):
	bgColour = (255,255,255)
	#check for hover state
	mousepos = pygame.mouse.get_pos()
	if pygame.Rect(pos, (80, 80)).collidepoint(mousepos):
		bgColour = (84,84,255)
		#detect mouse click
		if pygame.mouse.get_pressed()[0]:
			inGame.setInPlanet(True)
			inGame.setPlanetName(name)
	font = pygame.font.Font('freesansbold.ttf', 25) 
	text = font.render(name, True, bgColour) 
	textRect = text.get_rect()  
	textRect.center = (pos[0] + 40, pos[1] - 40) 
	screen.blit(text, textRect)
	pygame.draw.rect(screen, bgColour, (pos[0], pos[1], 80, 80))
	screen.blit(image, pos)