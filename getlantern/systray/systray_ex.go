package systray

import (
	"container/list"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/JyCyunMe/go-i18n/i18n"
)

type MenuItemEx struct {
	Item       *MenuItem
	Parent     *MenuItemEx
	Children   *list.List
	Callback   func(menuItemEx *MenuItemEx)
	I18nConfig *I18nConfig
}

type I18nConfig struct {
	TitleID           string
	TooltipID         string
	TitleFormat       string
	TooltipFormat     string
	TitleData         *i18n.Data
	TooltipData       *i18n.Data
	TitleNotAsTooltip bool
	Callback          func(i18nConfig *I18nConfig) string
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

// NilCallback 空回调
func NilCallback(menuItem *MenuItemEx) {
	//log.Infoln("clicked %s, id: %d", menuItem.Item.GetTitle(), menuItem.Item.GetId())
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

///

//func getI18nFormatted(i18nConfig *I18nConfig) (title string, tooltip string) {
func getI18nFormatted(i18nConfig *I18nConfig) (title string, tooltip string) {
	//title = i18n.T(i18nConfig.TitleID)
	title = i18n.TData("", i18nConfig.TitleID, i18nConfig.TitleData)
	if len(i18nConfig.TitleFormat) > 0 {
		if strings.Contains(i18nConfig.TitleFormat, "%s") {
			title = fmt.Sprintf(i18nConfig.TitleFormat, title)
		} else {
			title += i18nConfig.TitleFormat
		}
	}
	if len(i18nConfig.TooltipID) > 0 {
		//tooltip = i18n.T(i18nConfig.TooltipID)
		tooltip = i18n.TData("", i18nConfig.TooltipID, i18nConfig.TooltipData)
		if len(i18nConfig.TooltipFormat) > 0 {
			if strings.Contains(i18nConfig.TooltipFormat, "%s") {
				tooltip = fmt.Sprintf(i18nConfig.TooltipFormat, tooltip)
			} else {
				tooltip += i18nConfig.TooltipFormat
			}
		}
	} else if !i18nConfig.TitleNotAsTooltip {
		tooltip = title
	}
	return
}

// SwitchLanguage 切换语言
func (mie *MenuItemEx) SwitchLanguage() {
	if mie.I18nConfig != nil {
		title, tooltip := getI18nFormatted(mie.I18nConfig)
		mie.SetTitle(title)
		mie.SetTooltip(tooltip)
	}
}

// SwitchLanguageWithChildren 切换语言
func (mie *MenuItemEx) SwitchLanguageWithChildren() {
	mie.SwitchLanguage()
	for e := mie.Children.Front(); e != nil; e = e.Next() {
		child := e.Value.(*MenuItemEx)
		child.SwitchLanguageWithChildren()
	}
}

func (mie *MenuItemEx) setI18nConfig(i18nConfig *I18nConfig) (menuItemEx *MenuItemEx) {
	mie.I18nConfig = i18nConfig
	return mie
}

// AddMenuItemExI18n 添加增强版菜单项（同级）
func (mie *MenuItemEx) AddMenuItemExI18n(i18nConfig *I18nConfig, f func(menuItem *MenuItemEx)) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return mie.AddMenuItemEx(title, tooltip, f).setI18nConfig(i18nConfig)
}

// AddMenuItemExBindI18n 添加增强版菜单项（同级）并绑定到引用对象
func (mie *MenuItemEx) AddMenuItemExBindI18n(i18nConfig *I18nConfig, f func(menuItem *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return mie.AddMenuItemExBind(title, tooltip, f, v).setI18nConfig(i18nConfig)
}

// AddMenuItemCheckboxExI18n 添加增强版勾选框菜单项（同级）
func (mie *MenuItemEx) AddMenuItemCheckboxExI18n(i18nConfig *I18nConfig, isChecked bool, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return mie.AddMenuItemCheckboxEx(title, tooltip, isChecked, f).setI18nConfig(i18nConfig)
}

// AddMenuItemCheckboxExBindI18n 添加增强版菜单项并绑定到引用对象
func (mie *MenuItemEx) AddMenuItemCheckboxExBindI18n(i18nConfig *I18nConfig, isChecked bool, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return mie.AddMenuItemCheckboxExBind(title, tooltip, isChecked, f, v).setI18nConfig(i18nConfig)
}

// AddMainMenuItemExI18n 添加增强版主菜单项
func AddMainMenuItemExI18n(i18nConfig *I18nConfig, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return AddMainMenuItemEx(title, tooltip, f).setI18nConfig(i18nConfig)
}

// AddMainMenuItemExBindI18n 添加增强版主菜单项
func AddMainMenuItemExBindI18n(i18nConfig *I18nConfig, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	*v = *AddMainMenuItemExI18n(i18nConfig, f)
	return v.setI18nConfig(i18nConfig)
}

// AddSubMenuItemExI18n 添加增强版子菜单项
func (mie *MenuItemEx) AddSubMenuItemExI18n(i18nConfig *I18nConfig, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return mie.AddSubMenuItemEx(title, tooltip, f).setI18nConfig(i18nConfig)
}

// AddSubMenuItemExBindI18n 添加增强版子菜单项并绑定到引用对象
func (mie *MenuItemEx) AddSubMenuItemExBindI18n(i18nConfig *I18nConfig, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return mie.AddSubMenuItemExBind(title, tooltip, f, v).setI18nConfig(i18nConfig)
}

// AddSubMenuItemCheckboxExI18n 添加增强版勾选框子菜单项
func (mie *MenuItemEx) AddSubMenuItemCheckboxExI18n(i18nConfig *I18nConfig, isChecked bool, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return mie.AddSubMenuItemCheckboxEx(title, tooltip, isChecked, f).setI18nConfig(i18nConfig)
}

// AddSubMenuItemCheckboxExBindI18n 添加增强版勾选框子菜单项并绑定到引用对象
func (mie *MenuItemEx) AddSubMenuItemCheckboxExBindI18n(i18nConfig *I18nConfig, isChecked bool, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return mie.AddSubMenuItemCheckboxExBind(title, tooltip, isChecked, f, v).setI18nConfig(i18nConfig)
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
func (menuItemEx *MenuItemEx) UncheckFull() *MenuItemEx {
	for e := menuItemEx.Children.Front(); e != nil; e = e.Next() {
		e.Value.(*MenuItemEx).UncheckFull()
	}
	menuItemEx.Uncheck()
	return menuItemEx
}

// SetIcon sets the icon of a menu item. Only works on macOS and Windows.
// iconBytes should be the content of .ico/.jpg/.png
func (menuItemEx *MenuItemEx) SetIcon(iconBytes []byte) *MenuItemEx {
	menuItemEx.Item.SetIcon(iconBytes)
	return menuItemEx
}

// SetTemplateIcon sets the icon of a menu item as a template icon (on macOS). On Windows, it
// falls back to the regular icon bytes and on Linux it does nothing.
// templateIconBytes and regularIconBytes should be the content of .ico for windows and
// .ico/.jpg/.png for other platforms.
func (menuItemEx *MenuItemEx) SetTemplateIcon(templateIconBytes []byte, regularIconBytes []byte) *MenuItemEx {
	menuItemEx.Item.SetTemplateIcon(templateIconBytes, regularIconBytes)
	return menuItemEx
}

// SetTitle set the text to display on a menu item
func (menuItemEx *MenuItemEx) SetTitle(title string) *MenuItemEx {
	menuItemEx.Item.SetTitle(title)
	return menuItemEx
}

// SetTooltip set the tooltip to show when mouse hover
func (menuItemEx *MenuItemEx) SetTooltip(tooltip string) *MenuItemEx {
	menuItemEx.Item.SetTooltip(tooltip)
	return menuItemEx
}

// Disabled checks if the menu item is disabled
func (menuItemEx *MenuItemEx) Disabled() bool {
	return menuItemEx.Item.Disabled()
}

// Enable a menu item regardless if it's previously enabled or not
func (menuItemEx *MenuItemEx) Enable() *MenuItemEx {
	menuItemEx.Item.Enable()
	return menuItemEx
}

// Disable a menu item regardless if it's previously disabled or not
func (menuItemEx *MenuItemEx) Disable() *MenuItemEx {
	menuItemEx.Item.Disable()
	return menuItemEx
}

// Hide hides a menu item
func (menuItemEx *MenuItemEx) Hide() *MenuItemEx {
	menuItemEx.Item.Hide()
	return menuItemEx
}

// Show shows a previously hidden menu item
func (menuItemEx *MenuItemEx) Show() *MenuItemEx {
	menuItemEx.Item.Show()
	return menuItemEx
}

// Checked returns if the menu item has a check mark
func (menuItemEx *MenuItemEx) Checked() bool {
	return menuItemEx.Item.Checked()
}

// Check a menu item regardless if it's previously checked or not
func (menuItemEx *MenuItemEx) Check() *MenuItemEx {
	menuItemEx.Item.Check()
	return menuItemEx
}

// Uncheck a menu item regardless if it's previously unchecked or not
func (menuItemEx *MenuItemEx) Uncheck() *MenuItemEx {
	menuItemEx.Item.Uncheck()
	return menuItemEx
}

// Get ID of a menu item
func (menuItemEx *MenuItemEx) GetId() uint32 {
	return menuItemEx.Item.GetId()
}

// Get title of a menu item
func (menuItemEx *MenuItemEx) GetTitle() string {
	return menuItemEx.Item.GetTitle()
}

// Delete a menu item with children
func (menuItemEx *MenuItemEx) Delete() {
	menuItemEx.ClearChildren()
	menuItemEx.Hide()
}

func (menuItemEx *MenuItemEx) ClearChildren() *MenuItemEx {
	if menuItemEx.Children.Len() > 0 {
		lChild := menuItemEx.Children
		var next *list.Element
		for e := lChild.Front(); e != nil; e = next {
			next = e.Next()
			child := lChild.Remove(e).(*MenuItemEx)
			child.ClearChildren()
			child.Hide()
		}
	}
	menuItemEx.unsetSubMenu()
	return menuItemEx
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

func (menuItemEx *MenuItemEx) unsetSubMenu() *MenuItemEx {
	item := menuItemEx.Item
	_, err := wt.convertToNormalMenu(uint32(item.id))
	if err != nil {
		log.Errorf("Unable to unsetSubMenu: %v", err)
		return menuItemEx
	}
	return menuItemEx
}

func (menuItemEx *MenuItemEx) setSubMenu() *MenuItemEx {
	item := menuItemEx.Item
	_, err := wt.convertToSubMenu(uint32(item.id))
	if err != nil {
		log.Errorf("Unable to setSubMenu: %v", err)
		return menuItemEx
	}
	return menuItemEx
}
