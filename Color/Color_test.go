package Color_test

import (
    "testing"
    "bitbucket.org/prvn30/polstalker/Color"
)

func Test_Black_1(t *testing.T) {
    expected:="\x1b[30mpolru\x1b[39;49m"
    actual:=Color.Black("polru",true)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}

func Test_Black_2(t *testing.T) {
    expected:="polru"
    actual:=Color.Black("polru",false)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}

func Test_Red_1(t *testing.T) {
    expected:="\x1b[31mpolru\x1b[39;49m"
    actual:=Color.Red("polru",true)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}

func Test_Red_2(t *testing.T) {
    expected:="polru"
    actual:=Color.Red("polru",false)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}

func Test_Green_1(t *testing.T) {
    expected:="\x1b[32mpolru\x1b[39;49m"
    actual:=Color.Green("polru",true)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}

func Test_Green_2(t *testing.T) {
    expected:="polru"
    actual:=Color.Green("polru",false)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}

func Test_Yellow_1(t *testing.T) {
    expected:="\x1b[33mpolru\x1b[39;49m"
    actual:=Color.Yellow("polru",true)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}

func Test_Yellow_2(t *testing.T) {
    expected:="polru"
    actual:=Color.Yellow("polru",false)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}

func Test_Blue_1(t *testing.T) {
    expected:="\x1b[34mpolru\x1b[39;49m"
    actual:=Color.Blue("polru",true)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}

func Test_Blue_2(t *testing.T) {
    expected:="polru"
    actual:=Color.Blue("polru",false)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}

func Test_Magenta_1(t *testing.T) {
    expected:="\x1b[35mpolru\x1b[39;49m"
    actual:=Color.Magenta("polru",true)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}

func Test_Magenta_2(t *testing.T) {
    expected:="polru"
    actual:=Color.Magenta("polru",false)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}


func Test_Cyan_1(t *testing.T) {
    expected:="\x1b[36mpolru\x1b[39;49m"
    actual:=Color.Cyan("polru",true)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}

func Test_Cyan_2(t *testing.T) {
    expected:="polru"
    actual:=Color.Cyan("polru",false)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}

func Test_White_1(t *testing.T) {
    expected:="\x1b[37mpolru\x1b[39;49m"
    actual:=Color.White("polru",true)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}

func Test_White_2(t *testing.T) {
    expected:="polru"
    actual:=Color.White("polru",false)
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got: '%s'",expected,actual)
    }
}
