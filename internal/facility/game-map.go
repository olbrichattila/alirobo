package facility

import (
	"alibabarobotgame/internal/defaultconfig"
	"alibabarobotgame/internal/messages"
)

// initGameMap creates the entire underfund facility structure
func (f *fac) initGameMap() {
	collectPasswordFromEmployeeHint := messages.CollectPasswordFromEmployeeHint()
	collectPasswordFromComputerHint := messages.CollectPasswordFromComputerHint()
	bossRoomHintTest := messages.BossRoomHintText()

	f.undergroundFacility = facility{
		levels: []level{
			level{
				deepness: 0,
				rooms: []room{
					{bg: f.images.BossImage, floor: 250, ceil: 0, width: 210, leftWall: true, render: f.renderBossRoom, hintText: bossRoomHintTest},
					{bg: f.images.WallImage, floor: 250, ceil: 0, width: 390, roomLabel: "<< Boss room"},
					{bg: f.images.WallImage, floor: 250, ceil: 5, width: 380, passages: []passage{{pos: 30}}, rightWall: true},
					{width: 100},
					{bg: f.images.WallImage, floor: 260, ceil: 10, width: 600, rightWall: true, leftWall: true},
					{width: 200},
					{bg: f.images.Office2Image, floor: 200, ceil: 0, width: 320, leftWall: true, rightWall: true, passages: []passage{{pos: 200}}, render: f.renderOfficeFunctionComputeOffice, hintText: collectPasswordFromEmployeeHint},
					{width: 340},
					{bg: f.images.OfficeImage, floor: 300, ceil: 50, width: 610, leftWall: true, render: f.renderApsaraDBOffice, hintText: collectPasswordFromComputerHint},
					{bg: f.images.WallImage, floor: 300, ceil: 75, width: 380, passages: []passage{{pos: 30}}},
					{bg: f.images.Office2Image, floor: 300, ceil: 100, width: 320, rightWall: true, passages: []passage{{pos: 80}}, render: f.renderCdnOffice, hintText: collectPasswordFromEmployeeHint},
				},
			},
			level{
				deepness: 350,
				rooms: []room{
					{bg: f.images.WallImage, floor: 150, ceil: 0, width: 230, leftWall: true},
					{bg: f.images.WallImage, floor: 160, ceil: 5, width: 400, upperPassages: []passage{{pos: 40}}},
					{bg: f.images.WallImage, floor: 160, ceil: 10, width: 630, rightWall: true, passages: []passage{{pos: 250}}},
					{width: 50},
					{bg: f.images.OfficeImage, floor: 325, ceil: 0, width: 610, leftWall: true, passages: []passage{{pos: 50}}, render: f.renderOfficeEcsOffice, hintText: collectPasswordFromEmployeeHint},
					{bg: f.images.WallImage, floor: 300, ceil: 0, width: 610, upperPassages: []passage{{pos: 30}}},
					{bg: f.images.WallImage, floor: 300, ceil: 0, width: 310},
					{bg: f.images.Office3Image, floor: 310, ceil: 0, width: 799, rightWall: false, render: f.renderServerlessComputingOffice, passages: []passage{{pos: 500}}},
					{bg: f.images.WallImage, floor: 300, ceil: 0, width: 310, rightWall: true, upperPassages: []passage{{pos: 50}}},
				},
			},
			level{
				deepness: 800,
				rooms: []room{
					{bg: f.images.WallImage, floor: 200, ceil: 0, width: 230, leftWall: true},
					{bg: f.images.WallImage, floor: 200, ceil: 10, width: 400, passages: []passage{{pos: 80}}},
					{bg: f.images.WallImage, floor: 250, ceil: 0, width: 630, upperPassages: []passage{{pos: 60}}},
					{bg: f.images.WallImage, floor: 230, ceil: 0, width: 600},
					{bg: f.images.WallImage, floor: 230, ceil: 0, width: 600, rightWall: true, upperPassages: []passage{{pos: 550}}},
					{bg: f.images.Office2Image, floor: 200, ceil: 0, width: 300, leftWall: true, upperPassages: []passage{{pos: 30}}, render: f.renderObjectStorageServiceOffice, hintText: collectPasswordFromEmployeeHint},
					{bg: f.images.WallImage, floor: 200, ceil: 0, width: 700, rightWall: true, passages: []passage{{pos: 50}, {pos: 650}}},
				},
			},
			level{
				deepness: 1600,
				rooms: []room{
					{bg: f.images.WallImage, floor: 150, ceil: 0, width: 230, leftWall: true, passages: []passage{{pos: 50}}},
					{bg: f.images.WallImage, floor: 190, ceil: 4, width: 400},
					{bg: f.images.WallImage, floor: 265, ceil: 0, width: 380},
					{bg: f.images.WallImage, floor: 250, ceil: 0, width: 600, rightWall: true, upperPassages: []passage{{pos: 500}}},
					{width: 840},
					{bg: f.images.OfficeImage, floor: 325, ceil: 0, width: 610, rightWall: true, leftWall: true, upperPassages: []passage{{pos: 550}}, render: f.renderBlockStorageOffice, hintText: collectPasswordFromEmployeeHint},
					{width: 80},
					{bg: f.images.Office2Image, floor: 200, ceil: 0, width: 320, rightWall: true, leftWall: true, passages: []passage{{pos: 80}}, render: f.renderCloudBackupOffice, hintText: collectPasswordFromEmployeeHint},
				},
			},
			level{
				deepness: 2000,
				rooms: []room{
					{bg: f.images.WallImage, floor: 190, ceil: 0, width: 700, rightWall: true, leftWall: true, upperPassages: []passage{{pos: 600}}, passages: []passage{{pos: 50}, {pos: 650}}},
					{width: 60},
					{bg: f.images.WallImage, floor: 190, ceil: 0, width: 700, rightWall: true, leftWall: true, upperPassages: []passage{{pos: 600}}},
					{width: 1000},
					{bg: f.images.WallImage, floor: 190, ceil: 0, width: 650, leftWall: true, upperPassages: []passage{{pos: 50}}},
					{bg: f.images.WallImage, floor: 190, ceil: 4, width: 650, rightWall: true},
				},
			},
			level{
				deepness: 2400,
				rooms: []room{
					{bg: f.images.WallImage, floor: 190, ceil: 0, width: 550, rightWall: true, leftWall: true, passages: []passage{{pos: 50}}, upperPassages: []passage{{pos: 300}}},
					{width: 30},
					{bg: f.images.WallImage, floor: 190, ceil: 0, width: 650, leftWall: true, rightSafeDoor: true, roomType: defaultconfig.Ecs, upperPassages: []passage{{pos: 500}}},
					{bg: f.images.ServerRoomImage, floor: 190, ceil: 0, width: 400},
					{bg: f.images.WallImage, floor: 190, ceil: 0, width: 300, rightSafeDoor: true, roomType: defaultconfig.Cdn},
					{bg: f.images.ServerRoomImage, floor: 190, ceil: 0, width: 400},
					{bg: f.images.WallImage, floor: 190, ceil: 0, width: 300, rightSafeDoor: true, roomType: defaultconfig.BlockStorage},
					{bg: f.images.ServerRoomImage, floor: 190, ceil: 0, width: 400},
					{bg: f.images.WallImage, floor: 190, ceil: 0, width: 300, rightSafeDoor: true, roomType: defaultconfig.CloudBackup},
					{bg: f.images.ServerRoomImage, floor: 190, ceil: 0, width: 400, rightWall: false},
					{bg: f.images.WallImage, floor: 190, ceil: 0, width: 300, rightSafeDoor: true, roomType: defaultconfig.ServerlessComputing},
					{bg: f.images.ServerRoomImage, floor: 190, ceil: 0, width: 400, rightWall: true},
				},
			},
			level{
				deepness: 2700,
				rooms: []room{
					{bg: f.images.WallImage, floor: 190, ceil: 0, width: 650, leftWall: true, upperPassages: []passage{{pos: 600}}},
					{bg: f.images.WallImage, floor: 190, ceil: 0, width: 300, rightSafeDoor: true, roomType: defaultconfig.FunctionCompute},
					{bg: f.images.ServerRoomImage, floor: 190, ceil: 0, width: 400},
					{bg: f.images.WallImage, floor: 190, ceil: 0, width: 300, rightSafeDoor: true, roomType: defaultconfig.ObjectStorageService},
					{bg: f.images.ServerRoomImage, floor: 190, ceil: 0, width: 400},
					{bg: f.images.WallImage, floor: 190, ceil: 0, width: 600, rightSafeDoor: true, roomType: defaultconfig.ApsaraDB},
					{bg: f.images.ServerRoomImage, floor: 190, ceil: 0, width: 400, rightWall: true},
				},
			},
		},
	}
}
