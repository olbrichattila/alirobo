package facility

import (
	"alibabarobotgame/internal/defaultconfig"
)

func (f *fac) renderBossRoom() {
	f.inBossRoom = true
}

func (f *fac) renderOfficeEcsOffice() {
	f.renderCatchEmployeeToGetPassword(defaultconfig.Ecs)
}

func (f *fac) renderOfficeFunctionComputeOffice() {
	f.renderCatchEmployeeToGetPassword(defaultconfig.FunctionCompute)
}

func (f *fac) renderObjectStorageServiceOffice() {
	f.renderCatchEmployeeToGetPassword(defaultconfig.ObjectStorageService)
}

func (f *fac) renderBlockStorageOffice() {
	f.renderCatchEmployeeToGetPassword(defaultconfig.BlockStorage)
}

func (f *fac) renderCloudBackupOffice() {
	f.renderCatchEmployeeToGetPassword(defaultconfig.CloudBackup)
}

func (f *fac) renderCdnOffice() {
	f.renderCatchEmployeeToGetPassword(defaultconfig.Cdn)
}

func (f *fac) renderApsaraDBOffice() {
	if f.robotXinGrid > 2960 && f.robotXinGrid < 2990 {
		f.hoverOnBadgePickup = defaultconfig.ApsaraDB
	}
}

func (f *fac) renderServerlessComputingOffice() {
	if f.robotX > f.roomX+270 && f.robotX < f.roomX+340 {
		f.hoverOnBadgePickup = defaultconfig.ServerlessComputing
	}
}

func (f *fac) renderCatchEmployeeToGetPassword(eventType defaultconfig.AlibabaServiceType) {
	f.sprites.officeMan.SetY(f.roomBottom - 80)
	x := f.sprites.officeMan.GetX()
	if x >= f.roomXRight-50 || x >= defaultconfig.ScreenW-50 {
		f.officeWalkerOffset = -2
		f.sprites.officeMan.SetCollection(0)
	}

	if x <= f.roomX || x <= 0 {
		f.sprites.officeMan.SetCollection(1)
		f.officeWalkerOffset = +2
	}

	f.sprites.officeMan.SetX(x + f.officeWalkerOffset)
	f.sprites.officeMan.Render(f.screen)

	if f.robotX > x-10 && f.robotX < x+10 {
		f.hoverOnBadgePickup = eventType
	}

}
