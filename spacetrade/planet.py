import pygame
import gui
def drawPlanet(screen, inGame, planetsData, itemsData, inventoryManager):
	#title
	font = pygame.font.Font('freesansbold.ttf', 25) 
	text = font.render("IMT - Interplanetary Market Terminal", True, (255,255,255)) 
	textRect = text.get_rect()  
	textRect.topleft = (20, 20)
	screen.blit(text, textRect)
	planetData = planetsData[inGame.planetName]
	#slot 1
	drawItemCard(screen, planetData[0], (100, 200), itemsData, inventoryManager)
	#slot 2
	drawItemCard(screen, planetData[1], (280, 200), itemsData, inventoryManager)
	#slot 3
	drawItemCard(screen, planetData[2], (460, 200), itemsData, inventoryManager)
	#balance
	gui.button(screen, "$"+str(inventoryManager.bankBalance), (640 - 80, 30))
def drawItemCard(screen, name, pos, itemsData, inventoryManager):
	bgColour = (255,255,255)
	#check for hover state
	mousepos = pygame.mouse.get_pos()
	if pygame.Rect(pos, (80, 80)).collidepoint(mousepos):
		bgColour = (84,84,255)
	#bg
	pygame.draw.rect(screen, bgColour, (pos[0], pos[1], 80, 80))
	#name
	font = pygame.font.Font('freesansbold.ttf', 25) 
	text = font.render(name, True, bgColour) 
	textRect = text.get_rect()  
	textRect.center = (pos[0] + 40, pos[1] - 40) 
	screen.blit(text, textRect)
	#price
	font = pygame.font.Font('freesansbold.ttf', 20) 
	text = font.render("$"+str(itemsData[name]), True, (0,0,0)) 
	textRect = text.get_rect()  
	textRect.center = (pos[0] + 40, pos[1] + 40) 
	screen.blit(text, textRect)
	#quantity owned
	font = pygame.font.Font('freesansbold.ttf', 20) 
	text = font.render("Q: "+str(inventoryManager.inventoryData[name]), True, (255,255,255)) 
	textRect = text.get_rect()  
	textRect.center = (pos[0] + 40, pos[1] + 200) 
	screen.blit(text, textRect)
	#buttons
	if(gui.greenButton(screen, "Buy", (pos[0]+10, pos[1]+115))):
		buy(name, inventoryManager, itemsData)
	if(gui.redButton(screen, "Sell", (pos[0]+70, pos[1]+115))):
		sell(name, inventoryManager, itemsData)

def buy(name, inventoryManager, itemsData):
	if (inventoryManager.bankBalance - itemsData[name] >= 0):
		inventoryManager.setBalance(inventoryManager.bankBalance - itemsData[name])
		inventoryManager.incItem(name)

def sell(name, inventoryManager, itemsData):
	if (inventoryManager.inventoryData[name] - 1 >= 0):
		inventoryManager.setBalance(inventoryManager.bankBalance + itemsData[name])
		inventoryManager.decItem(name)

def drawSpaceStation(screen, inGame, inventoryManager):
	itemsData = {
		"Market Speed+": 1000.0,
		"Market Speed-": 1000.0,
		"I AM RICH": 100000.0
	}
	#title
	font = pygame.font.Font('freesansbold.ttf', 25) 
	text = font.render("Space Station Kiosk", True, (255,255,255)) 
	textRect = text.get_rect()  
	textRect.topleft = (20, 20)
	screen.blit(text, textRect)
	#slot 1
	drawItemCardCustom(screen, "Market Speed+", (100, 200), itemsData, inventoryManager, inGame)
	#slot 2
	drawItemCardCustom(screen, "Market Speed-", (280, 200), itemsData, inventoryManager, inGame)
	#slot 3
	drawItemCard(screen, "I AM RICH", (460, 200), itemsData, inventoryManager)
	#balance
	gui.button(screen, "$"+str(inventoryManager.bankBalance), (640 - 80, 30))

def drawItemCardCustom(screen, name, pos, itemsData, inventoryManager, inGame):
	bgColour = (255,255,255)
	#check for hover state
	mousepos = pygame.mouse.get_pos()
	if pygame.Rect(pos, (80, 80)).collidepoint(mousepos):
		bgColour = (84,84,255)
	#bg
	pygame.draw.rect(screen, bgColour, (pos[0], pos[1], 80, 80))
	#name
	font = pygame.font.Font('freesansbold.ttf', 25) 
	text = font.render(name, True, bgColour) 
	textRect = text.get_rect()  
	textRect.center = (pos[0] + 40, pos[1] - 40) 
	screen.blit(text, textRect)
	#price
	font = pygame.font.Font('freesansbold.ttf', 20) 
	text = font.render("$"+str(itemsData[name]), True, (0,0,0)) 
	textRect = text.get_rect()  
	textRect.center = (pos[0] + 40, pos[1] + 40) 
	screen.blit(text, textRect)
	#buttons
	if(gui.greenButton(screen, "Buy", (pos[0]+40, pos[1]+115))):
		buyCustom(name, inventoryManager, itemsData, inGame)

def buyCustom(name, inventoryManager, itemsData, inGame):
	if (inventoryManager.bankBalance - itemsData[name] >= 0):
		inventoryManager.setBalance(inventoryManager.bankBalance - itemsData[name])
		if name == "Market Speed+":
			inGame.setCounterMax(inGame.counterMax-100)
		elif name == "Market Speed-":
			inGame.setCounterMax(inGame.counterMax+100)