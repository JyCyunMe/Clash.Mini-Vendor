package systray

import (
	"github.com/JyCyunMe/go-i18n/i18n"
)

type MenuItemI18nConfig struct {
	TitleConfig       *i18n.I18nConfig
	TooltipConfig     *i18n.I18nConfig
	TitleNotAsTooltip bool
	//CallbackData        *i18n.CallbackData
}

type I18nConfig struct {
	TitleID           string
	TooltipID         string
	TitleFormat       string
	TooltipFormat     string
	TitleData         *i18n.Data
	TooltipData       *i18n.Data
	TitleNotAsTooltip bool
}

func NewI18nConfig(i18nConfig I18nConfig) (menuItemI18nConfig *MenuItemI18nConfig) {
	menuItemI18nConfig = new(MenuItemI18nConfig)
	menuItemI18nConfig.TitleConfig = &i18n.I18nConfig{
		Id:     i18nConfig.TitleID,
		Format: i18nConfig.TitleFormat,
		Data:   i18nConfig.TitleData,
	}
	menuItemI18nConfig.TooltipConfig = &i18n.I18nConfig{
		Id:     i18nConfig.TooltipID,
		Format: i18nConfig.TooltipFormat,
		Data:   i18nConfig.TooltipData,
	}
	menuItemI18nConfig.TitleNotAsTooltip = i18nConfig.TitleNotAsTooltip
	return
}

func getI18nFormatted(i18nConfig *MenuItemI18nConfig) (title string, tooltip string) {
	if i18nConfig.TitleConfig == nil {
		title = ""
	} else {
		title = i18n.GetI18nFormatted(i18nConfig.TitleConfig)
	}
	if i18nConfig.TooltipConfig != nil {
		tooltip = i18n.GetI18nFormatted(i18nConfig.TooltipConfig)
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
	mie.ForChildrenLoop(true, func(_ int, child *MenuItemEx) {
		child.SwitchLanguageWithChildren()
	})
}

func (mie *MenuItemEx) setI18nConfig(i18nConfig *MenuItemI18nConfig) (menuItemEx *MenuItemEx) {
	mie.I18nConfig = i18nConfig
	return mie
}

// AddMenuItemExI18n 添加增强版菜单项（同级）
func (mie *MenuItemEx) AddMenuItemExI18n(i18nConfig *MenuItemI18nConfig, f func(menuItem *MenuItemEx)) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return mie.AddMenuItemEx(title, tooltip, f).setI18nConfig(i18nConfig)
}

// AddMenuItemExBindI18n 添加增强版菜单项（同级）并绑定到引用对象
func (mie *MenuItemEx) AddMenuItemExBindI18n(i18nConfig *MenuItemI18nConfig, f func(menuItem *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	*v = *mie.AddMenuItemExI18n(i18nConfig, f)
	return v.setI18nConfig(i18nConfig)
}

// AddMenuItemCheckboxExI18n 添加增强版勾选框菜单项（同级）
func (mie *MenuItemEx) AddMenuItemCheckboxExI18n(i18nConfig *MenuItemI18nConfig, isChecked bool, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return mie.AddMenuItemCheckboxEx(title, tooltip, isChecked, f).setI18nConfig(i18nConfig)
}

// AddMenuItemCheckboxExBindI18n 添加增强版菜单项并绑定到引用对象
func (mie *MenuItemEx) AddMenuItemCheckboxExBindI18n(i18nConfig *MenuItemI18nConfig, isChecked bool, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return mie.AddMenuItemCheckboxExBind(title, tooltip, isChecked, f, v).setI18nConfig(i18nConfig)
}

// AddMainMenuItemExI18n 添加增强版主菜单项
func AddMainMenuItemExI18n(i18nConfig *MenuItemI18nConfig, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return AddMainMenuItemEx(title, tooltip, f).setI18nConfig(i18nConfig)
}

// AddMainMenuItemExBindI18n 添加增强版主菜单项
func AddMainMenuItemExBindI18n(i18nConfig *MenuItemI18nConfig, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	*v = *AddMainMenuItemExI18n(i18nConfig, f)
	return v.setI18nConfig(i18nConfig)
}

// AddSubMenuItemExI18n 添加增强版子菜单项
func (mie *MenuItemEx) AddSubMenuItemExI18n(i18nConfig *MenuItemI18nConfig, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return mie.AddSubMenuItemEx(title, tooltip, f).setI18nConfig(i18nConfig)
}

// AddSubMenuItemExBindI18n 添加增强版子菜单项并绑定到引用对象
func (mie *MenuItemEx) AddSubMenuItemExBindI18n(i18nConfig *MenuItemI18nConfig, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	*v = *mie.AddSubMenuItemExI18n(i18nConfig, f)
	return v.setI18nConfig(i18nConfig)
}

// AddSubMenuItemCheckboxExI18n 添加增强版勾选框子菜单项
func (mie *MenuItemEx) AddSubMenuItemCheckboxExI18n(i18nConfig *MenuItemI18nConfig, isChecked bool, f func(menuItemEx *MenuItemEx)) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return mie.AddSubMenuItemCheckboxEx(title, tooltip, isChecked, f).setI18nConfig(i18nConfig)
}

// AddSubMenuItemCheckboxExBindI18n 添加增强版勾选框子菜单项并绑定到引用对象
func (mie *MenuItemEx) AddSubMenuItemCheckboxExBindI18n(i18nConfig *MenuItemI18nConfig, isChecked bool, f func(menuItemEx *MenuItemEx), v *MenuItemEx) (menuItemEx *MenuItemEx) {
	title, tooltip := getI18nFormatted(i18nConfig)
	return mie.AddSubMenuItemCheckboxExBind(title, tooltip, isChecked, f, v).setI18nConfig(i18nConfig)
}
