// Package messages returns different hard coded messages
package messages

func NotShapeShifted() []string {
	return []string{
		"You have lost a life", "",
		"You should shape shift to trick",
		"the employee, but you did not, therefore",
		"the security team was informed and you were",
		"removed.",
		"", "Press enter to continue",
	}
}

func CollectPasswordFromEmployeeHint() []string {
	return []string{
		"As a shape shifter robot, first you need to",
		"shift to a human with pressing S, then",
		"when the employee comes to you, press space quickly",
		"to obtain entry code, let's say you tell him, I am",
		"an engineer and I left my pass card at home, can you",
		"borrow your pass code for today?",
	}
}

func CollectPasswordFromComputerHint() []string {
	return []string{
		"Someone left the computer unlocked",
		"and logged into password manager.",
		"Find the computer in the room and press space",
		"to collect the passcode to enter a server room",
	}
}

func BossRoomHintText() []string {
	return []string{
		"Your boss given you, as an ethical hacker",
		"shape shifter robot to collect entry codes for server rooms",
		"with social engineering",
	}
}

func LooserText() []string {
	return []string{
		"Whoops, it was not successful this time", "",
		"Press enter to restart the game",
	}
}

func WinnerText() []string {
	return []string{
		"Congratulation, you have completed the task, and you entered all server rooms with social",
		"with obtaining entry codes to all server rooms using social engineering",
		"",
		" - enter your name to be on the score board",
	}
}

/*
Game Introduction:
In this game, you play as a shape-shifting ethical hacker robot on a mission assigned by your boss. Your objective is to use social engineering tactics to infiltrate eight different server rooms located in a secret underground facility.
Navigate the maze of tunnels and ladders using the arrow keys. Your first goal is to locate all the office rooms.
If an office has an employee moving around, you must first shape-shift into a human form by pressing the S key. Be aware that if you move afterward, you will revert to your original robot form. Wait for the employee to come close, then press the Shift key to interact. This represents a social engineering attempt where you claim to be an engineer who forgot their access badge. If successful, you’ll receive a new badge, and a corresponding server room icon will appear, indicating the type of server room the badge grants access to.
Some offices won’t have employees but will contain a computer on a desk. Approach the computer and press the Space bar to collect a password. This simulates the ethical hacker discovering an unattended, unlocked machine and retrieving credentials from it.
Once you’ve collected all the badges, they will begin to flash, signaling that you are now ready to locate and access the server rooms on levels -5 and -6.
Each server room requires the correct badge to enter. The type of server room is indicated by an icon displayed near the door, representing different Alibaba services. You must recognize the correct icon and use the corresponding badge. You have three attempts per room—after that, you’ll lose a life.
Successfully entering a server room allows you to continue to the next one.
After infiltrating all server rooms, return to the starting room, where your boss is waiting. If all server rooms have been accessed, your mission is complete. This final step symbolizes reporting your findings to your boss so they can initiate appropriate security training.
At any time, if you're inside a room and unsure what to do, press the H key to receive a hint—provided the room has special rules or tasks.
*/
func GameIntro() []string {
	return []string{
		"In this game, you play as a shape-shifting ethical hacker robot",
		"on a mission assigned by your boss. Your objective is to use",
		"social engineering tactics to infiltrate eight different server rooms",
		"located in a secret underground facility.",
		"",
		"Navigate the maze of tunnels and ladders using the arrow keys.",
		"Your first goal is to locate all the office rooms.",
		"If an office has an employee moving around, you must first shape-shift",
		"into a human form by pressing the S key. Be aware that if you move",
		"afterward, you will revert to your original robot form.",
		"Wait for the employee to come close, then press the Shift key to interact.",
		"This represents a social engineering attempt where you claim to be an engineer",
		"who forgot their access badge. If successful, you’ll receive a new badge, and a",
		"corresponding server room icon will appear, indicating the type of server room",
		"the badge grants access to.",
		"",
		"Some offices won’t have employees but will contain a computer on a desk.",
		"Approach the computer and press the Space bar to collect a password.",
		"This simulates the ethical hacker discovering an unattended, unlocked",
		"machine and retrieving credentials from it.",
		"Once you’ve collected all the badges, they will begin to flash, signaling",
		"that you are now ready to locate and access the server rooms on levels -5 and -6.",
		"",
		"Each server room requires the correct badge to enter.",
		"The type of server room is indicated by an icon displayed near the door,",
		"representing different Alibaba services. You must recognize the correct icon",
		"and use the corresponding badge. You have three attempts per room—after that,",
		"you’ll lose a life.",
		"Successfully entering a server room allows you to continue to the next one.",
		"After infiltrating all server rooms, return to the starting room, where your",
		"boss is waiting. If all server rooms have been accessed, your mission is complete.",
		"This final step symbolizes reporting your findings to your boss so they can initiate",
		"appropriate security training.",
		"",
		"At any time, if you're inside a room and unsure what to do,",
		"press the H key to receive a hint—provided the room has special rules or tasks.",
		"",
		"You have 20 minutes to complete the task.",
	}
}
