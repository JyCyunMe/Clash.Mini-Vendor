package systray

import (
	"container/list"
	"os"
	"runtime"
)

type MenuItemEx struct {
	Item       *MenuItem
	Parent     *MenuItemEx
	Children   *list.List
	Callback   func(menuItemEx *MenuItemEx)
	I18nConfig *MenuItemI18nConfig
	ExtraData  interface{}
}

var (
	MenuList []*MenuItemEx
)

// RunEx SystrayEx入口 须在init()调用
func RunEx(onReady func(), onExit func()) {
	// use it on init
	go func() {
		runtime.LockOSThread()
		Run(onReady, func() {
			onExit()
			os.Exit(1)
		})
		runtime.UnlockOSThread()
	}()
}

// ShowMenu 弹出托盘菜单
func ShowMenu() error {
	return wt.showMenu()
}

// NilCallback 空回调
func NilCallback(menuItem *MenuItemEx) {
	//log.Infoln("clicked %s, id: %d", menuItem.Item.GetTitle(), menuItem.Item.GetId())
}

// ClickFunc 点击回调
func ClickFunc(clickFunc func()) func() {
	return func() {
		go clickFunc()
	}
}

// SetLeftClickFunc 设置左键点击回调
func SetLeftClickFunc(leftClickFunc func()) {
	wt.lClickFunc = leftClickFunc
}

// SetRightClickFunc 设置右键点击回调
func SetRightClickFunc(rightClickFunc func()) {
	wt.rClickFunc = rightClickFunc
}

// AddMenuItemEx 添加增强版菜单项（同级）
func (mie *MenuItemEx) AddMenuItemEx(title string, tooltip string, f func(menuItem *MenuItemEx)) (menuItemEx *MenuItemEx) {
	menuItemEx = getSubMenuItemEx(mie.Parent.Item, title, tooltip, f)
	menuItemEx.Parent = mie.Parent
	mie.Parent.Children.PushBack(menuItemEx)
	return
}

// AddMenuItemExBind 添加增强版菜单项（同级）并绑定到引用对象
func (mie *MenuItemEx) AddMenuItemExBind(title string, tooltip string, f func(menuItem *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	menuItemEx = getSubMenuItemEx(mie.Parent.Item, title, tooltip, f)
	menuItemEx.Parent = mie.Parent
	mie.Parent.Children.PushBack(menuItemEx)
	*v = *menuItemEx
	return
}

// AddMenuItemCheckboxEx 添加增强版勾选框菜单项（同级）
func (mie *MenuItemEx) AddMenuItemCheckboxEx(title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	menuItemEx = getSubMenuItemCheckboxEx(mie.Parent.Item, title, tooltip, isChecked, f)
	menuItemEx.Parent = mie.Parent
	mie.Parent.Children.PushBack(menuItemEx)
	return
}

// AddMenuItemCheckboxExBind 添加增强版菜单项并绑定到引用对象
func (mie *MenuItemEx) AddMenuItemCheckboxExBind(title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	menuItemEx = getSubMenuItemCheckboxEx(mie.Parent.Item, title, tooltip, isChecked, f)
	menuItemEx.Parent = mie.Parent
	mie.Parent.Children.PushBack(menuItemEx)
	*v = *menuItemEx
	return
}

// AddMainMenuItemEx 添加增强版主菜单项
func AddMainMenuItemEx(title string, tooltip string, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	menuItemEx = getMenuItemEx(title, tooltip, f)
	MenuList = append(MenuList, menuItemEx)
	return
}

// AddMainMenuItemExBind 添加增强版主菜单项并绑定到引用对象
func AddMainMenuItemExBind(title string, tooltip string, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	menuItemEx = AddMainMenuItemEx(title, tooltip, f)
	*v = *menuItemEx
	return
}

// AddSubMenuItemEx 添加增强版子菜单项
func (mie *MenuItemEx) AddSubMenuItemEx(title string, tooltip string, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	//subMenuItemEx := getMenuItemEx(title, tooltip, f)
	//mie.Children = append(mie.Children, subMenuItemEx)
	//return mie
	menuItemEx = getSubMenuItemEx(mie.Item, title, tooltip, f)
	menuItemEx.Parent = mie
	mie.Children.PushBack(menuItemEx)
	//mie.setSubMenu()
	return
}

// AddSubMenuItemExBind 添加增强版子菜单项并绑定到引用对象
func (mie *MenuItemEx) AddSubMenuItemExBind(title string, tooltip string, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	menuItemEx = mie.AddSubMenuItemEx(title, tooltip, f)
	*v = *menuItemEx
	return
}

// AddSubMenuItemCheckboxEx 添加增强版勾选框子菜单项
func (mie *MenuItemEx) AddSubMenuItemCheckboxEx(title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	//subMenuItemEx := getMenuItemEx(title, tooltip, f)
	//mie.Children = append(mie.Children, subMenuItemEx)
	//return mie
	menuItemEx = getSubMenuItemCheckboxEx(mie.Item, title, tooltip, isChecked, f)
	menuItemEx.Parent = mie
	mie.Children.PushBack(menuItemEx)
	//mie.setSubMenu()
	return
}

// AddSubMenuItemCheckboxExBind 添加增强版勾选框子菜单项并绑定到引用对象
func (mie *MenuItemEx) AddSubMenuItemCheckboxExBind(title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	menuItemEx = mie.AddSubMenuItemCheckboxEx(title, tooltip, isChecked, f)
	*v = *menuItemEx
	return
}

//// AddSeparator adds a separator bar to the menu
//func AddSeparator(mie *MenuItemEx) *MenuItemEx {
//	menuItemEx := &MenuItemEx{
//	}
//	addSeparator(menuItemEx.GetId())
//	//addSeparator(atomic.AddUint32(&currentID, 1))
//	return menuItemEx
//}

// SwitchCheckboxGroup 切换增强版勾选框菜单项组 设置指定项勾选与否，组内其他项相反
func SwitchCheckboxGroup(newValue *MenuItemEx, checked bool, values []*MenuItemEx) {
	for _, value := range values {
		if value.GetId() == newValue.GetId() {
			if checked {
				value.Check()
			} else {
				value.Uncheck()
			}
		} else {
			if checked {
				value.Uncheck()
			} else {
				value.Check()
			}
		}
	}
}

// SwitchCheckboxBrother 切换增强版勾选框菜单项组 设置指定项勾选与否，其他兄弟项相反
func (mie *MenuItemEx) SwitchCheckboxBrother(checked bool) {
	SwitchCheckboxGroupByList(mie, checked, mie.Parent.Children)
}

// SwitchCheckboxGroupByList 切换增强版勾选框菜单项组 设置指定项勾选与否，组内其他项相反
func SwitchCheckboxGroupByList(newValue *MenuItemEx, checked bool, values *list.List) {
	if values == nil || values.Len() == 0 {
		newValue.Checked()
	}
	for e := values.Front(); e != nil; e = e.Next() {
		value := e.Value.(*MenuItemEx)
		if value.GetId() == newValue.GetId() {
			if checked {
				value.Check()
			} else {
				value.UncheckFull()
			}
		} else {
			if checked {
				value.UncheckFull()
			} else {
				value.Check()
			}
		}
	}
}

// UncheckFull uncheck with children
func (mie *MenuItemEx) UncheckFull() *MenuItemEx {
	for e := mie.Children.Front(); e != nil; e = e.Next() {
		e.Value.(*MenuItemEx).UncheckFull()
	}
	mie.Uncheck()
	return mie
}

// SetIcon sets the icon of a menu item. Only works on macOS and Windows.
// iconBytes should be the content of .ico/.jpg/.png
func (mie *MenuItemEx) SetIcon(iconBytes []byte) *MenuItemEx {
	mie.Item.SetIcon(iconBytes)
	return mie
}

// SetTemplateIcon sets the icon of a menu item as a template icon (on macOS). On Windows, it
// falls back to the regular icon bytes and on Linux it does nothing.
// templateIconBytes and regularIconBytes should be the content of .ico for windows and
// .ico/.jpg/.png for other platforms.
func (mie *MenuItemEx) SetTemplateIcon(templateIconBytes []byte, regularIconBytes []byte) *MenuItemEx {
	mie.Item.SetTemplateIcon(templateIconBytes, regularIconBytes)
	return mie
}

// SetTitle set the text to display on a menu item
func (mie *MenuItemEx) SetTitle(title string) *MenuItemEx {
	mie.Item.SetTitle(title)
	return mie
}

// SetTooltip set the tooltip to show when mouse hover
func (mie *MenuItemEx) SetTooltip(tooltip string) *MenuItemEx {
	mie.Item.SetTooltip(tooltip)
	return mie
}

// Disabled checks if the menu item is disabled
func (mie *MenuItemEx) Disabled() bool {
	return mie.Item.Disabled()
}

// Enable a menu item regardless if it's previously enabled or not
func (mie *MenuItemEx) Enable() *MenuItemEx {
	mie.Item.Enable()
	return mie
}

// Disable a menu item regardless if it's previously disabled or not
func (mie *MenuItemEx) Disable() *MenuItemEx {
	mie.Item.Disable()
	return mie
}

// Hide hides a menu item
func (mie *MenuItemEx) Hide() *MenuItemEx {
	mie.Item.Hide()
	return mie
}

// Show shows a previously hidden menu item
func (mie *MenuItemEx) Show() *MenuItemEx {
	mie.Item.Show()
	return mie
}

// Checked returns if the menu item has a check mark
func (mie *MenuItemEx) Checked() bool {
	return mie.Item.Checked()
}

// Check a menu item regardless if it's previously checked or not
func (mie *MenuItemEx) Check() *MenuItemEx {
	mie.Item.Check()
	return mie
}

// Uncheck a menu item regardless if it's previously unchecked or not
func (mie *MenuItemEx) Uncheck() *MenuItemEx {
	mie.Item.Uncheck()
	return mie
}

// GetId Get ID of a menu item
func (mie *MenuItemEx) GetId() uint32 {
	return mie.Item.GetId()
}

// GetTitle Get title of a menu item
func (mie *MenuItemEx) GetTitle() string {
	return mie.Item.GetTitle()
}

// GetTooltip Get tooltip of a menu item
func (mie *MenuItemEx) GetTooltip() string {
	return mie.Item.tooltip
}

// GetExtraData Get Extra Data of a menu item
func (mie *MenuItemEx) GetExtraData() interface{} {
	return mie.ExtraData
}

// SetExtraData Get Extra Data of a menu item
func (mie *MenuItemEx) SetExtraData(extraData interface{}) {
	mie.ExtraData = extraData
}

// Delete a menu item with children
func (mie *MenuItemEx) Delete() {
	mie.ClearChildren()
	mie.Hide()
}

func (mie *MenuItemEx) ClearChildren() *MenuItemEx {
	if mie.Children.Len() > 0 {
		lChild := mie.Children
		var next *list.Element
		for e := lChild.Front(); e != nil; e = next {
			next = e.Next()
			child := lChild.Remove(e).(*MenuItemEx)
			child.ClearChildren()
			child.Hide()
		}
	}
	mie.unsetSubMenu()
	return mie
}

func getMenuItemEx(title string, tooltip string, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	menuItem := AddMenuItem(title, tooltip)
	menuItemEx = &MenuItemEx{
		Item:     menuItem,
		Children: list.New(),
	}
	menuItem.setExObj(menuItemEx)
	menuItemEx.Callback = func(e *MenuItemEx) {
		go f(menuItemEx)
	}
	return menuItemEx
}

func getSubMenuItemEx(menuItem *MenuItem, title string, tooltip string, f func(menuItemEx *MenuItemEx)) (subMenuItemEx *MenuItemEx) {
	subMenuItem := menuItem.AddSubMenuItem(title, tooltip)
	subMenuItemEx = &MenuItemEx{
		Item:     subMenuItem,
		Children: list.New(),
	}
	subMenuItem.setExObj(subMenuItemEx)
	subMenuItemEx.Callback = func(e *MenuItemEx) {
		go f(subMenuItemEx)
	}
	return subMenuItemEx
}

func getSubMenuItemCheckboxEx(menuItem *MenuItem, title string, tooltip string, isChecked bool, f func(menuItemEx *MenuItemEx)) (subMenuItemEx *MenuItemEx) {
	subMenuItem := menuItem.AddSubMenuItemCheckbox(title, tooltip, isChecked)
	subMenuItemEx = &MenuItemEx{
		Item:     subMenuItem,
		Children: list.New(),
	}
	subMenuItem.setExObj(subMenuItemEx)
	subMenuItemEx.Callback = func(e *MenuItemEx) {
		go f(subMenuItemEx)
	}
	return subMenuItemEx
}

func (mie *MenuItemEx) unsetSubMenu() *MenuItemEx {
	item := mie.Item
	_, err := wt.convertToNormalMenu(uint32(item.id))
	if err != nil {
		log.Errorf("Unable to unsetSubMenu: %v", err)
		return mie
	}
	return mie
}

func (mie *MenuItemEx) setSubMenu() *MenuItemEx {
	item := mie.Item
	_, err := wt.convertToSubMenu(uint32(item.id))
	if err != nil {
		log.Errorf("Unable to setSubMenu: %v", err)
		return mie
	}
	return mie
}
