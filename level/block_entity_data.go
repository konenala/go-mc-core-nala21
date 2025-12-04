package level

// BlockEntityData is a generic interface for all block entity data
type BlockEntityData interface {
	// isBlockEntityData is a placeholder method to satisfy the interface.
	isBlockEntityData()
}

// BaseBlockEntity holds common fields for all block entities.
type BaseBlockEntity struct {
	ID string `nbt:"id"`
	X  int32  `nbt:"x"`
	Y  int32  `nbt:"y"`
	Z  int32  `nbt:"z"`
	// KeepPacked tells the unmarshaller to not unpack this struct
	KeepPacked bool `nbt:"keep_packed"`
}

func (b BaseBlockEntity) isBlockEntityData() {}

// SignEntity represents the data for a sign.
// Note: This is a simplified representation. Newer versions use a more complex structure
// with front_text and back_text objects containing messages, color, etc.
type SignEntity struct {
	BaseBlockEntity
	Text1   string `nbt:"Text1"`
	Text2   string `nbt:"Text2"`
	Text3   string `nbt:"Text3"`
	Text4   string `nbt:"Text4"`
	Color   string `nbt:"Color"`
	Glowing bool   `nbt:"GlowingText"`
	IsWaxed bool   `nbt:"is_waxed"`
}

func (s SignEntity) isBlockEntityData() {}

// ChestEntity represents the data for a chest.
// This is a placeholder for now.
type ChestEntity struct {
	BaseBlockEntity
	// Items []... // To be implemented
}

func (c ChestEntity) isBlockEntityData() {}
