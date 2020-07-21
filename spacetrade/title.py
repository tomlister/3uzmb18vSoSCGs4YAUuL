import gui
class titleScreen:
	def __init__(self):
		self.textAnimationPosition = 0
		self.buttonOpacity = 0
	def drawTitleScreen(self, screen, font, inGame):
		if self.textAnimationPosition != 50:
			self.textAnimationPosition = self.textAnimationPosition + 2
		text = font.render('Commercium', True, (255,255,255)) 
		textRect = text.get_rect()  
		textRect.center = (640 / 2, (480 / 2) - self.textAnimationPosition) 
		screen.blit(text, textRect)
		if(gui.button(screen, "Start", (640 / 2, (480 / 2) + 50))):
			inGame.setInGame(True)
