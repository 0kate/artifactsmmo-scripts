package characters

import (
	"github.com/0kate/artifactsmmo-scripts/internal/items"
	"github.com/0kate/artifactsmmo-scripts/internal/shared"
)

type Level = int
type XP = int
type Gold = int
type Speed = int
type HP = int
type Attack = int
type Damage = int
type Resistance = int

type Character struct {
	name    string
	skin    Skin
	level   Level
	xp      XP
	maxXP   XP
	totalXP XP
	gold    Gold
	speed   Speed

	miningLevel Level
	miningXP    XP
	miningMaxXP XP

	woodcuttingLevel Level
	woodcuttingXP    XP
	woodcuttingMaxXP XP

	fishingLevel Level
	fishingXP    XP
	fishingMaxXP XP

	weaponCraftingLevel Level
	weaponCraftingXP    XP
	weaponCraftingMaxXP XP

	gearCraftingLevel Level
	gearCraftingXP    XP
	gearCraftingMaxXP XP

	jewelyCraftingLevel Level
	jewelyCraftingXP    XP
	jewelyCraftingMaxXP XP

	cookingLevel Level
	cookingXP    XP
	cookingMaxXP XP

	hp             HP
	haste          int
	criticalStrike int
	stamina        int

	attackFire  Attack
	attackEarth Attack
	attackWater Attack
	attackAir   Attack

	damageFire  Damage
	damageEarth Damage
	damageWater Damage
	damageAir   Damage

	resistanceFire  Resistance
	resistanceEarth Resistance
	resistanceWater Resistance
	resistanceAir   Resistance

	position *shared.Position

	cooldown Cooldown

	weaponSlot              items.Code
	shieldSlot              items.Code
	helmetSlot              items.Code
	bodyArmorSlot           items.Code
	legArmorSlot            items.Code
	bootsSlot               items.Code
	ring1Slot               items.Code
	ring2Slot               items.Code
	amuletSlot              items.Code
	artifact1Slot           items.Code
	artifact2Slot           items.Code
	artifact3Slot           items.Code
	consumable1Slot         items.Code
	consumable1SlotQuantity int
	consumable2Slot         items.Code
	consumable2SlotQuantity int

	task         string
	taskType     TaskType
	taskProgress TaskProgress
	taskTotal    TaskTotal

	inventoryMaxItems int
	inventory         []InventorySlot
}

func NewCharacter(name string) *Character {
	return &Character{name: name}
}

func (c *Character) Name() string {
	return c.name
}

func (c *Character) Skin() Skin {
	return c.skin
}

func (c *Character) Level() Level {
	return c.level
}

func (c *Character) XP() XP {
	return c.xp
}

func (c *Character) Inventory() []InventorySlot {
	return c.inventory
}
