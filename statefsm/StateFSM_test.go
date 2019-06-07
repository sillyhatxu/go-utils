package statefsm

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 电风扇
type ElectricFan struct {
	*FSM
}

// 实例化电风扇
func NewElectricFan(initState FSMState) *ElectricFan {
	return &ElectricFan{
		FSM: NewFSM(initState),
	}
}

var (
	PowerOff   = FSMState("关闭")
	FirstGear  = FSMState("1档")
	SecondGear = FSMState("2档")
	ThirdGear  = FSMState("3档")

	PowerOffEvent   = FSMEvent{EventName: "按下关闭按钮", EventState: PowerOff}
	FirstGearEvent  = FSMEvent{EventName: "按下1档按钮", EventState: FirstGear}
	SecondGearEvent = FSMEvent{EventName: "按下2档按钮", EventState: SecondGear}
	ThirdGearEvent  = FSMEvent{EventName: "按下3档按钮", EventState: ThirdGear}
)

func TestFSM(t *testing.T) {
	efan := NewElectricFan(PowerOff)                                                                        // 初始状态是关闭的
	efan.AddHandler(PowerOff, []FSMEvent{PowerOffEvent, FirstGearEvent, SecondGearEvent, ThirdGearEvent})   // 关闭状态
	efan.AddHandler(FirstGear, []FSMEvent{PowerOffEvent, FirstGearEvent, SecondGearEvent, ThirdGearEvent})  // 1档状态
	efan.AddHandler(SecondGear, []FSMEvent{PowerOffEvent, FirstGearEvent, SecondGearEvent, ThirdGearEvent}) // 2档状态
	efan.AddHandler(ThirdGear, []FSMEvent{PowerOffEvent, FirstGearEvent, SecondGearEvent, ThirdGearEvent})  // 3档状态
	// 开始测试状态变化
	err := efan.Call(ThirdGearEvent, FSMHandler(func() error {
		log.Println("电风扇开启3档，发型被吹乱了！")
		return nil
	})) // 按下3档按钮
	assert.Nil(t, err)
	err = efan.Call(FirstGearEvent, FSMHandler(func() error {
		log.Println("电风扇开启1档，微风徐来！")
		return nil
	})) // 按下1档按钮
	assert.Nil(t, err)
	err = efan.Call(PowerOffEvent, FSMHandler(func() error {
		log.Println("电风扇已关闭")
		return nil
	})) // 按下关闭按钮
	assert.Nil(t, err)
	err = efan.Call(SecondGearEvent, FSMHandler(func() error {
		log.Println("电风扇开启2档，凉飕飕！")
		return nil
	})) // 按下2档按钮
	assert.Nil(t, err)
	err = efan.Call(PowerOffEvent, FSMHandler(func() error {
		log.Println("电风扇已关闭")
		return nil
	})) // 按下关闭按钮
	assert.Nil(t, err)
}

func TestFSM2(t *testing.T) {
	efan := NewElectricFan(PowerOff)                                                       // 初始状态是关闭的
	efan.AddHandler(PowerOff, []FSMEvent{FirstGearEvent})                                  // 关闭状态
	efan.AddHandler(FirstGear, []FSMEvent{PowerOffEvent, SecondGearEvent})                 // 1档状态
	efan.AddHandler(SecondGear, []FSMEvent{PowerOffEvent, FirstGearEvent, ThirdGearEvent}) // 2档状态
	efan.AddHandler(ThirdGear, []FSMEvent{PowerOffEvent, SecondGearEvent})                 // 3档状态
	// 开始测试状态变化
	err := efan.Call(ThirdGearEvent, FSMHandler(func() error {
		log.Println("电风扇开启3档，发型被吹乱了！")
		return nil
	})) // 按下3档按钮
	assert.NotNil(t, err)
	assert.EqualValues(t, efan.state, PowerOff)

	err = efan.Call(FirstGearEvent, FSMHandler(func() error {
		log.Println("电风扇开启1档，微风徐来！")
		return nil
	})) // 按下1档按钮
	assert.Nil(t, err)

	err = efan.Call(ThirdGearEvent, FSMHandler(func() error {
		log.Println("电风扇开启3档，发型被吹乱了！")
		return nil
	})) // 按下3档按钮
	assert.NotNil(t, err)
	assert.EqualValues(t, efan.state, FirstGearEvent.EventState)

	err = efan.Call(SecondGearEvent, FSMHandler(func() error {
		log.Println("电风扇开启2档，凉飕飕！")
		return nil
	})) // 按下2档按钮
	assert.Nil(t, err)

	err = efan.Call(ThirdGearEvent, FSMHandler(func() error {
		log.Println("电风扇开启3档，发型被吹乱了！")
		return nil
	})) // 按下3档按钮
	assert.Nil(t, err)

	err = efan.Call(PowerOffEvent, FSMHandler(func() error {
		log.Println("电风扇已关闭")
		return nil
	})) // 按下关闭按钮
	assert.Nil(t, err)

	err = efan.Call(PowerOffEvent, FSMHandler(func() error {
		log.Println("电风扇已关闭")
		return nil
	})) // 按下关闭按钮
	assert.NotNil(t, err)
}
