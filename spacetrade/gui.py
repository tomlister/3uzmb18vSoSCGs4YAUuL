import pygame
def button(screen, content, pos):
	font = pygame.font.Font('freesansbold.ttf', 25) 
	text = font.render(content, True, (0,0,0), (255,255,255)) 
	textRect = text.get_rect()
	textRect.center = pos
	mousepos = pygame.mouse.get_pos()
	#check for button hover state
	if textRect.collidepoint(mousepos):
		text = font.render(content, True, (255,255,255), (84,84,255)) 
		screen.blit(text, textRect)
		#detect mouse click
		if pygame.mouse.get_pressed()[0]:
			return True
		else:
			return False
	screen.blit(text, textRect)

def greenButton(screen, content, pos):
	font = pygame.font.Font('freesansbold.ttf', 25) 
	text = font.render(content, True, (0,0,0), (84,255,84)) 
	textRect = text.get_rect()
	textRect.center = pos
	mousepos = pygame.mouse.get_pos()
	#check for button hover state
	if textRect.collidepoint(mousepos):
		text = font.render(content, True, (0,0,0), (255,255,255)) 
		screen.blit(text, textRect)
		#detect mouse click
		if pygame.mouse.get_pressed()[0]:
			return True
		else:
			return False
	screen.blit(text, textRect)

def redButton(screen, content, pos):
	font = pygame.font.Font('freesansbold.ttf', 25) 
	text = font.render(content, True, (0,0,0), (255,84,84)) 
	textRect = text.get_rect()
	textRect.center = pos
	mousepos = pygame.mouse.get_pos()
	#check for button hover state
	if textRect.collidepoint(mousepos):
		text = font.render(content, True, (0,0,0), (255,255,255)) 
		screen.blit(text, textRect)
		#detect mouse click
		if pygame.mouse.get_pressed()[0]:
			return True
		else:
			return False
	screen.blit(text, textRect)