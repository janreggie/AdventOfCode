package aoc2015

import (
	"testing"

	"github.com/antonholmquist/jason"
	"github.com/janreggie/AdventOfCode/tools"
	"github.com/stretchr/testify/assert"
)

func Test_extractSum(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		value string // to be marshalled using NewValueFromBytes
		want  int64
	}{
		// TODO: Add test cases.
		{"int64", `43`, 43},
		{"array", `[1,2,3]`, 6},
		{"object", `{"a":2,"b":4}`, 6},
		{"nested array", `[[[3]]]`, 3},
		{"nested object", `{"a":{"b":4},"c":-1}`, 3},
		{"object with array value", `{"a":[-1,1]}`, 0},
		{"array of int and object", `[-1,{"a":1}]`, 0},
	}
	for _, tt := range tests {
		value, err := jason.NewValueFromBytes([]byte(tt.value))
		assert.NoError(err, tt.name)
		assert.Equal(tt.want, extractSum(value), tt.name)
	}
}

func Test_extractSumButRed(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		value string // to be marshalled using NewValueFromBytes
		want  int64
	}{
		// TODO: Add test cases.
		{"int64", `43`, 43},
		{"array but red object", `[1,{"c":"red","b":2},3]`, 4},
		{"object with red", `{"d":"red","e":[1,2,3,4],"f":5}`, 0},
		{"array with red string", `[1,"red",5]`, 6},
	}
	for _, tt := range tests {
		value, err := jason.NewValueFromBytes([]byte(tt.value))
		assert.NoError(err, tt.name)
		assert.Equal(tt.want, extractSumButRed(value), tt.name)
	}
}
func TestDay12(t *testing.T) {
	assert := assert.New(t)
	testCases := []tools.TestCase{
		{Details: "Y2015D12 my input",
			Input:
			// This is supposed to be convoluted.
			// I truncated my input in 40-wide bytes.
			`{"e":[[{"e":86,"c":23,"a":{"a":[120,169,` +
				`"green","red","orange"],"b":"red"},"g":"` +
				`yellow","b":["yellow"],"d":"red","f":-19` +
				`},{"e":-47,"a":[2],"d":{"a":"violet"},"c` +
				`":"green","h":"orange","b":{"e":59,"a":"` +
				`yellow","d":"green","c":47,"h":"red","b"` +
				`:"blue","g":"orange","f":["violet",43,16` +
				`8,78]},"g":"orange","f":[{"e":[82,-41,2,` +
				`"red","violet","orange","yellow"],"c":"g` +
				`reen","a":77,"g":"orange","b":147,"d":49` +
				`,"f":"blue"},-1,142,136,["green","red",1` +
				`66,-21],"blue","orange",{"a":38}]},"oran` +
				`ge","yellow"],"green",-22,[37,[4,-40,["r` +
				`ed","yellow",["yellow",177,"red","blue",` +
				`139,[55,13,"yellow","violet",-21,140,"ye` +
				`llow",117],"blue","blue",106],"blue",{"a` +
				`":23}],183,92,"orange","green"],"orange"` +
				`],-5],"c":[{"e":{"e":-13,"c":-11,"a":{"a` +
				`":49,"b":189},"g":144,"b":186,"d":{"e":[` +
				`146,[32,"violet","red","orange",-22],"bl` +
				`ue","violet",57,{"e":12,"a":"red","d":37` +
				`,"c":-13,"h":"green","b":-27,"g":"orange` +
				`","f":"orange","i":"red"},56,-1,"red",-2` +
				`5],"c":-14,"a":[["orange","green","green` +
				`","red",-25],-16,104,177,"red"],"g":"red` +
				`","b":"blue","d":2,"f":"green"},"f":[{"e` +
				`":{"c":-15,"a":"green","b":144,"d":-32},` +
				`"c":"yellow","a":["blue","blue"],"b":"ye` +
				`llow","d":135,"f":"violet"}]},"a":{"e":"` +
				`blue","a":[145,128,"orange","violet",23,` +
				`["orange",[78,"yellow","orange","orange"` +
				`,"orange","green",122,-35,"blue"],159,11` +
				`4]],"d":"yellow","c":{"e":[100,"green",{` +
				`"e":"blue","a":36,"d":84,"j":"orange","c` +
				`":"blue","h":118,"b":85,"g":"violet","f"` +
				`:"blue","i":27},"blue","blue"],"a":"red"` +
				`,"d":"blue","j":[159,22,"violet","orange` +
				`","blue","orange","blue",186,175,{"e":29` +
				`,"a":150,"d":"green","c":129,"h":138,"b"` +
				`:-29,"g":7,"f":"red","i":"violet"}],"c":` +
				`"yellow","h":["blue",-9,41,{"e":144,"c":` +
				`"violet","a":161,"b":116,"d":126,"f":197` +
				`},173,123,50,"red",154],"b":-23,"g":"ora` +
				`nge","f":"green","i":"red"},"h":26,"b":3` +
				`0,"g":106,"f":{"c":{"e":-28,"c":["red","` +
				`violet","blue","orange","yellow","violet` +
				`",-30,134,0],"a":-37,"b":41,"d":143},"a"` +
				`:121,"b":"yellow","d":[198,"yellow"]}},"` +
				`d":{"e":-2,"a":"green","d":{"a":0,"b":"b` +
				`lue"},"j":"violet","c":[{"e":"violet","a` +
				`":"orange","d":"violet","c":-36,"h":68,"` +
				`b":195,"g":41,"f":63},11,-20,{"e":[100,4` +
				`3,183],"a":"yellow","d":["orange","yello` +
				`w","violet","yellow",143,162,-23,168,145` +
				`,-33],"j":80,"c":"orange","h":{"a":"oran` +
				`ge"},"b":"violet","g":-21,"f":-38,"i":"g` +
				`reen"},{"e":"red","a":81,"d":3,"c":"blue` +
				`","h":[167,"blue","yellow",135,64,116,13` +
				`4,14,160],"b":"red","g":"blue","f":"gree` +
				`n"},["red",[139,"orange","blue","blue","` +
				`orange","blue"],"orange","red",113,26,"y` +
				`ellow","yellow",85]],"h":[[117],"yellow"` +
				`,{"e":187,"c":"orange","a":-22,"b":["ora` +
				`nge"],"d":68},154,["green","orange","blu` +
				`e",87,"green","orange",46,"violet"],{"a"` +
				`:{"e":178,"c":"blue","a":-2,"g":99,"b":"` +
				`orange","d":-34,"f":"green"}},"blue"],"b` +
				`":"violet","g":[13,"yellow",43,"orange"]` +
				`,"f":{"e":-40,"a":"blue","d":50,"c":"vio` +
				`let","h":{"e":"orange","a":[-27,159,"vio` +
				`let","red",130,83,"red","violet",-27],"d` +
				`":-42,"j":"blue","c":9,"h":{"e":"yellow"` +
				`,"a":"blue","d":177,"c":156,"h":"violet"` +
				`,"b":0,"g":"orange","f":-43,"i":"orange"` +
				`},"b":"blue","g":[114,"blue",-4,"red","r` +
				`ed","yellow","green"],"f":81,"i":"red"},` +
				`"b":73,"g":176,"f":"red"},"i":{"e":["ora` +
				`nge",-16,["violet",63,"blue",-40,119,22,` +
				`"orange","orange","violet","green"],-30,` +
				`{"e":"yellow","c":172,"a":"green","g":"y` +
				`ellow","b":"green","d":"yellow","f":"red` +
				`"},190,28,{"e":"orange","c":"red","a":"o` +
				`range","b":51,"d":-31,"f":136},"orange",` +
				`82],"a":"blue","d":"red","c":["blue",81,` +
				`"blue","blue","green",191,"blue","blue",` +
				`1],"h":{"e":"orange","c":{"e":"blue","a"` +
				`:"green","d":55,"j":-36,"c":13,"h":"yell` +
				`ow","b":162,"g":82,"f":"red","i":"red"},` +
				`"a":"violet","b":"yellow","d":-24,"f":19` +
				`0},"b":[{"e":"green","a":"green","d":-49` +
				`,"j":"green","c":"orange","h":43,"b":"or` +
				`ange","g":35,"f":"violet","i":"blue"},"y` +
				`ellow",32,"yellow"],"g":"blue","f":0}},"` +
				`c":"blue","h":69,"b":166,"g":[[88,["blue` +
				`",21,"yellow","violet"]],["red",169],"re` +
				`d"],"f":176},{"e":["yellow",88,164,{"e":` +
				`"red","c":"yellow","a":[20,"blue","viole` +
				`t"],"g":-31,"b":80,"d":"yellow","f":"gre` +
				`en"},{"e":"orange","c":"green","a":149,"` +
				`b":"orange","d":-46,"f":[160,83,"orange"` +
				`,"red",177,-11]},"green",[156,"red",{"c"` +
				`:7,"a":[152,107,130],"b":{"c":"yellow","` +
				`a":114,"b":38,"d":"blue"},"d":"orange"},` +
				`{"a":49,"b":-34},34,-32,"green"]],"a":13` +
				`4,"d":-12,"c":[-33,{"e":102,"a":"red","d` +
				`":{"e":"orange","a":"green","d":43,"c":"` +
				`violet","h":{"a":"violet"},"b":-24,"g":"` +
				`blue","f":"blue","i":[68,"blue"]},"j":-2` +
				`8,"c":87,"h":"violet","b":122,"g":"viole` +
				`t","f":"green","i":"violet"}],"h":"viole` +
				`t","b":["orange",-14,{"e":37,"a":86,"d":` +
				`{"e":186,"a":-25,"d":71,"c":"orange","h"` +
				`:86,"b":113,"g":27,"f":"green","i":"yell` +
				`ow"},"c":16,"h":"orange","b":"green","g"` +
				`:["yellow","yellow",["orange","violet","` +
				`violet","green",58,"orange"],["orange",1` +
				`31,"red","blue","orange",183,82,"orange"` +
				`,"yellow","blue"],49,"orange","violet","` +
				`violet","yellow","orange"],"f":"yellow"}` +
				`,[141,"violet","red",-24,18,103,88,169,7` +
				`5],["yellow",["green",55,92,"yellow","or` +
				`ange",135,{"e":"green","c":141,"a":-11,"` +
				`b":129,"d":"orange","f":"green"},"violet` +
				`"],14,[{"a":130},["red","violet",182,"bl` +
				`ue",149,"orange",-25,"blue","blue"],61,-` +
				`18,"orange",14,{"e":"red","a":135,"d":"y` +
				`ellow","j":"red","c":23,"h":89,"b":82,"g` +
				`":"orange","f":"red","i":72},"red"],["re` +
				`d","blue","red",78,134,53,160],-20,98,{"` +
				`e":[198,69,168,145,-29,"red","orange","o` +
				`range","yellow"],"a":"violet","d":"green` +
				`","j":184,"c":[7,99,186,"blue"],"h":10,"` +
				`b":"blue","g":166,"f":"yellow","i":177}]` +
				`,"green",{"e":"red","a":"green","d":-8,"` +
				`j":-47,"c":{"e":"orange","c":74,"a":"yel` +
				`low","b":"orange","d":34,"f":124},"h":15` +
				`2,"b":"red","g":"yellow","f":161,"i":["b` +
				`lue","red","orange","orange","orange",-3` +
				`8,"orange","red"]},"yellow","blue",75],"` +
				`g":[67,["orange",109,114,32,"green","gre` +
				`en","yellow",["yellow","orange",-40,["gr` +
				`een","orange","yellow",187,3,"yellow","v` +
				`iolet","orange",195,"yellow"],"blue","ye` +
				`llow","blue"]],[32,{"c":43,"a":"red","b"` +
				`:"blue","d":25}],{"c":"red","a":24,"b":1` +
				`39},{"e":"orange","a":153,"d":43,"c":143` +
				`,"h":["violet","yellow","green",159,165,` +
				`{"a":"blue","b":"violet"}],"b":150,"g":[` +
				`"red","red","orange",[138,"green"]],"f":` +
				`"green"},"violet",-12,"red",["violet","g` +
				`reen"]],"f":{"e":{"e":39,"a":"orange","d` +
				`":{"e":[11],"c":"violet","a":"orange","g` +
				`":49,"b":"red","d":0,"f":{"a":45}},"c":"` +
				`violet","h":-30,"b":[93,-1,"red",[39,"re` +
				`d","green"],"green",[154,"blue","orange"` +
				`,147,"orange","yellow"],106,["green",71,` +
				`-9],-37],"g":"blue","f":"blue","i":176},` +
				`"c":98,"a":164,"b":["violet","orange",[8` +
				`5,{"e":"yellow","a":113,"d":176,"c":"gre` +
				`en","h":"violet","b":"orange","g":166,"f` +
				`":"green","i":"green"},[145,"green",-7,"` +
				`violet"]],{"e":["green","green"],"a":"re` +
				`d","d":79,"j":"blue","c":-41,"h":"yellow` +
				`","b":"violet","g":"blue","f":148,"i":20` +
				`}],"d":12,"f":[-34,"yellow"]}},{"e":{"e"` +
				`:[{"a":159},["red",-46,{"a":11},166,116,` +
				`{"e":-39,"c":-24,"a":194,"b":27,"d":91},` +
				`-37,85,["green"],61],"yellow",{"a":135},` +
				`["orange","orange",128,"green",-20,97,{"` +
				`e":108,"c":195,"a":"blue","g":51,"b":"gr` +
				`een","d":"violet","f":28}]],"a":"green",` +
				`"d":164,"c":"violet","h":"orange","b":"g` +
				`reen","g":{"e":"red","c":151,"a":"violet` +
				`","b":46,"d":"yellow","f":["red"]},"f":-` +
				`35,"i":"orange"},"a":{"e":7,"a":{"a":-28` +
				`},"d":151,"c":{"c":147,"a":165,"b":[-5,[` +
				`"violet","blue","orange","violet",3,"yel` +
				`low",86,"orange",197,51],6,156,43,94,"bl` +
				`ue",{"e":130,"c":"orange","a":-29,"b":89` +
				`,"d":-41,"f":"orange"},76]},"h":{"e":153` +
				`,"a":"violet","d":"yellow","j":"green","` +
				`c":{"c":"blue","a":"violet","b":113,"d":` +
				`"yellow"},"h":{"e":"blue","a":["yellow",` +
				`29,69],"d":"orange","c":"red","h":"green` +
				`","b":164,"g":"blue","f":{"c":16,"a":191` +
				`,"b":61}},"b":195,"g":"yellow","f":"gree` +
				`n","i":"red"},"b":82,"g":[51,-47,186,{"e` +
				`":4,"c":27,"a":60,"b":"orange","d":32,"f` +
				`":"violet"},{"e":"blue","a":72,"d":17,"j` +
				`":"blue","c":"red","h":0,"b":"yellow","g` +
				`":195,"f":["red","green",82,-31,"blue",-` +
				`24,"yellow","red","violet"],"i":43},["gr` +
				`een"]],"f":"blue"},"d":["violet",22,118]` +
				`,"j":-45,"c":174,"h":79,"b":180,"g":{"c"` +
				`:-7,"a":{"e":["blue","violet"],"c":"blue` +
				`","a":"violet","g":"red","b":"orange","d` +
				`":"orange","f":{"a":"violet","b":33}},"b` +
				`":183,"d":132},"f":["violet","violet","g` +
				`reen",[[[66,"violet","violet","green","g` +
				`reen"]],[181,"yellow",167,134,"orange",{` +
				`"e":"red","c":"violet","a":"violet","b":` +
				`107,"d":-19},{"e":0,"c":166,"a":"green",` +
				`"b":"blue","d":"red"}],19,-31,108]],"i":` +
				`["red","red",[159],[139,"blue",{"a":106}` +
				`,48,117,164,["blue",161,"green",174,"ora` +
				`nge",152,"red","orange",["red","yellow",` +
				`"blue",-43]],105,22,"green"],[110]]},153` +
				`,{"a":{"e":"orange","a":[{"e":"red","a":` +
				`"yellow","d":"green","c":"violet","h":{"` +
				`e":"blue","c":62,"a":148,"g":"violet","b` +
				`":6,"d":"yellow","f":-1},"b":"yellow","g` +
				`":-14,"f":58},"violet","yellow",{"e":4,"` +
				`a":"blue","d":{"e":"green","a":-49,"d":"` +
				`yellow","j":-39,"c":"orange","h":"red","` +
				`b":"blue","g":59,"f":"violet","i":46},"c` +
				`":195,"h":22,"b":160,"g":"orange","f":"o` +
				`range","i":38},"red",["orange","red","ye` +
				`llow",34,101,"yellow"],40,["orange",{"e"` +
				`:148,"c":"red","a":85,"g":62,"b":-13,"d"` +
				`:-25,"f":"orange"},-47,34,{"e":"violet",` +
				`"c":80,"a":"red","b":34,"d":100},58,185,` +
				`"yellow","orange",["yellow","green","vio` +
				`let",84,"blue","orange",13]],183],"d":"r` +
				`ed","j":[-2,"red","yellow",176,-24,140,"` +
				`blue","yellow",155,{"e":-28,"c":"violet"` +
				`,"a":{"a":"green","b":3},"b":"red","d":-` +
				`16}],"c":["violet",-16],"h":"blue","b":"` +
				`green","g":["yellow","yellow","yellow",4` +
				`4,"orange",50,36,{"e":"green","a":162,"d` +
				`":112,"c":166,"h":92,"b":31,"g":"blue","` +
				`f":-12}],"f":"yellow","i":["yellow","vio` +
				`let",[19,"red",["violet","violet",195]],` +
				`["yellow",[106],"red","orange","blue"],1` +
				`78,{"e":"violet","a":104,"d":"red","c":"` +
				`yellow","h":{"c":-47,"a":"blue","b":84},` +
				`"b":"green","g":0,"f":"blue","i":"red"},` +
				`105]}},{"e":7,"c":[156],"a":{"a":{"e":{"` +
				`e":"violet","a":["blue","yellow","orange` +
				`"],"d":"green","c":[159,"blue","violet",` +
				`"red",61,3],"h":"red","b":"green","g":13` +
				`2,"f":130},"c":92,"a":"orange","b":"gree` +
				`n","d":[186,122],"f":"yellow"}},"g":["ye` +
				`llow","violet"],"b":"violet","d":{"e":{"` +
				`c":"orange","a":127,"b":41,"d":[36,61,17` +
				`8,"yellow","green","red","violet",{"e":-` +
				`18,"a":"yellow","d":"red","c":0,"h":"yel` +
				`low","b":"yellow","g":6,"f":"yellow"},16` +
				`2]},"c":-6,"a":{"c":{"e":"red","c":78,"a` +
				`":"blue","b":91,"d":49,"f":14},"a":["gre` +
				`en"],"b":{"a":22},"d":{"a":"blue"}},"b":` +
				`154,"d":"orange","f":{"a":170}},"f":-33}` +
				`,"blue",82],"a":{"c":["green",["red","or` +
				`ange",{"e":-28,"a":{"a":"violet","b":110` +
				`},"d":[[174,140,72],191,"yellow",108,195` +
				`,{"a":"violet"},147,53],"c":"yellow","h"` +
				`:"green","b":"violet","g":"red","f":["bl` +
				`ue","orange","violet",[48,118],156,144,-` +
				`46,110,["orange","yellow","blue","red"],` +
				`149],"i":{"e":"orange","c":101,"a":{"e":` +
				`111,"a":"blue","d":"orange","j":"orange"` +
				`,"c":-40,"h":13,"b":"orange","g":"yellow` +
				`","f":32,"i":"yellow"},"g":{"e":"orange"` +
				`,"a":"blue","d":195,"j":81,"c":185,"h":2` +
				`0,"b":4,"g":"green","f":112,"i":147},"b"` +
				`:-22,"d":199,"f":"yellow"}},"yellow",19,` +
				`128,-3,27,["orange",{"e":-8,"c":156,"a":` +
				`"yellow","b":"red","d":20,"f":-37},[{"c"` +
				`:19,"a":"blue","b":150},"orange",-12,9]]` +
				`],[12,{"e":"blue","c":162,"a":["blue",18` +
				`4,"yellow","orange",{"a":"yellow","b":"g` +
				`reen"},88,-19,60,"yellow"],"g":"yellow",` +
				`"b":191,"d":-6,"f":"violet"}],{"c":"yell` +
				`ow","a":"orange","b":{"a":"violet","b":[` +
				`"orange","orange","violet",{"e":"red","a` +
				`":"red","d":163,"c":153,"h":"green","b":` +
				`6,"g":"blue","f":17,"i":63},163,[164,-41` +
				`,"violet","violet",126]]},"d":-38}],"a":` +
				`{"e":{"c":-1,"a":"orange","b":{"c":131,"` +
				`a":{"e":-11,"c":120,"a":"green","b":198,` +
				`"d":152,"f":37},"b":77,"d":{"e":8,"a":21` +
				`,"d":"blue","c":"yellow","h":"violet","b` +
				`":11,"g":"violet","f":{"e":148,"c":98,"a` +
				`":80,"b":78,"d":68}}},"d":"orange"},"a":` +
				`["violet",[-30,117],[78,31],74,197,"red"` +
				`,"orange",95],"d":"green","c":[96,"viole` +
				`t"],"h":{"e":{"c":"green","a":[76,16,125` +
				`,"green",15,"violet",130,60,"red"],"b":"` +
				`orange","d":-38},"a":71,"d":158,"j":-16,` +
				`"c":[["yellow","green",183,165,-28,4,102` +
				`],-20,"blue","violet",{"e":"yellow","c":` +
				`{"e":"orange","a":"yellow","d":"red","c"` +
				`:"orange","h":"orange","b":169,"g":"viol` +
				`et","f":48},"a":99,"b":["blue",-1,"blue"` +
				`],"d":104,"f":20},83],"h":"green","b":[-` +
				`14,[28],"yellow",[93,"blue",-24,160,35,2` +
				`5,-32,"green"],{"e":[51,"red",64,"red","` +
				`blue",-16,31,146,"blue","yellow"],"c":12` +
				`2,"a":"orange","b":"yellow","d":{"c":53,` +
				`"a":179,"b":"blue","d":-44}},17,110],"g"` +
				`:151,"f":"orange","i":{"a":"violet","b":` +
				`{"a":[-24]}}},"b":["violet",{"e":{"e":91` +
				`,"a":{"e":"blue","a":-25,"d":70,"c":"gre` +
				`en","h":"violet","b":48,"g":"violet","f"` +
				`:"orange"},"d":"yellow","c":136,"h":90,"` +
				`b":{"e":26,"c":"green","a":"blue","g":"v` +
				`iolet","b":192,"d":198,"f":86},"g":"oran` +
				`ge","f":"blue"},"a":137,"d":47,"c":11,"h` +
				`":"yellow","b":"orange","g":"orange","f"` +
				`:{"e":"blue","a":45,"d":"violet","j":146` +
				`,"c":-38,"h":4,"b":157,"g":104,"f":-13,"` +
				`i":"yellow"}},{"e":191,"a":"blue","d":"g` +
				`reen","c":"blue","h":-19,"b":148,"g":"bl` +
				`ue","f":{"e":57,"c":"red","a":167,"b":[-` +
				`42,147,166,74,-32,"orange","violet","yel` +
				`low"],"d":"green","f":71}},"green",[184,` +
				`"yellow",[["yellow","yellow"],"green",{"` +
				`e":-24,"a":1,"d":44,"c":"yellow","h":"ye` +
				`llow","b":144,"g":"violet","f":"green","` +
				`i":-7},86,119,52,"orange",["red","red",1` +
				`8,"orange",192,116],120,109]],["violet",` +
				`-14,"violet",{"e":-29,"c":{"e":"green","` +
				`a":"red","d":"blue","j":68,"c":9,"h":"or` +
				`ange","b":25,"g":"red","f":10,"i":"green` +
				`"},"a":-15,"b":"blue","d":"violet","f":1` +
				`25},119,[127,"violet","green",39]]],"g":` +
				`"green","f":-14},"b":23,"d":{"c":{"c":98` +
				`,"a":"yellow","b":97},"a":-29,"b":{"a":1` +
				`92,"b":["violet","yellow",65,{"c":{"e":"` +
				`violet","c":"yellow","a":"violet","b":"b` +
				`lue","d":"orange"},"a":"red","b":176},19` +
				`2]},"d":"orange"}},"b":[{"e":"yellow","c` +
				`":45,"a":81,"b":["orange"],"d":"violet",` +
				`"f":[-3,"red",146,186,"orange","red","bl` +
				`ue",{"e":"green","c":22,"a":"yellow","b"` +
				`:"blue","d":-2,"f":"green"},0,180]},[[-3` +
				`6,["orange",[166],"violet"],{"c":86,"a":` +
				`[2,173,78,"violet","orange",["violet","y` +
				`ellow","blue",107,24,-1,"orange",13,"gre` +
				`en","violet"]],"b":"violet","d":107},100` +
				`,["yellow",-22,[177,69,144,84,159,"viole` +
				`t"],"green"],{"e":"green","a":78,"d":173` +
				`,"c":"blue","h":36,"b":[[-48,164,"red","` +
				`blue",45],["green","orange",23,15,110,49` +
				`,"blue"],"violet",0,192,53],"g":["blue",` +
				`"violet"],"f":"orange","i":[{"e":186,"c"` +
				`:"orange","a":"green","b":174,"d":"yello` +
				`w","f":46},"violet",188,"yellow",54,-6,"` +
				`blue","violet",0,{"e":166,"a":"yellow","` +
				`d":"red","j":"blue","c":"red","h":97,"b"` +
				`:"violet","g":32,"f":173,"i":95}]},{"a":` +
				`58,"b":"blue"},"green",["red",150,3,"ora` +
				`nge",32,106,[["blue"],118,{"c":178,"a":7` +
				`,"b":185,"d":"violet"},"red",164,"red",[` +
				`128,"red"],-44],{"e":"orange","a":"orang` +
				`e","d":130,"c":"yellow","h":"yellow","b"` +
				`:{"e":91,"a":161,"d":-44,"c":-45,"h":"bl` +
				`ue","b":"orange","g":122,"f":"orange"},"` +
				`g":"yellow","f":"blue"},139,{"a":97}]],"` +
				`orange",["orange",0,"blue","red",{"e":19` +
				`2,"c":92,"a":{"a":["red",164,"yellow",18` +
				`9,"blue",150,"green","violet",-35,33],"b` +
				`":"red"},"g":"yellow","b":{"a":"blue"},"` +
				`d":"red","f":-31},111,"yellow","red",[["` +
				`green","green",42,-47,[88,142,"blue",59,` +
				`-42,"violet"],"green"],"orange",["violet` +
				`","yellow","violet",198,94,44,"orange","` +
				`green","blue",26],"blue","violet"]],{"a"` +
				`:118}],{"e":-35,"a":{"e":["orange",-1,12` +
				`1,"red"],"c":"violet","a":[[127,-18,-4,[` +
				`-40,42,"violet",167,"orange",112,"orange` +
				`"],30,31,"violet",37],{"c":"blue","a":"g` +
				`reen","b":172},[141,154],146,"yellow"],"` +
				`b":"blue","d":-3},"d":"red","c":-17,"h":` +
				`-3,"b":["violet","yellow",19,"red",8,138` +
				`,37],"g":{"c":{"e":-5,"c":[-23,21,"green` +
				`",-3,"red"],"a":163,"b":"blue","d":"red"` +
				`,"f":["violet",136,"violet"]},"a":183,"b` +
				`":-36,"d":"violet"},"f":["green",["green` +
				`",{"e":"red","c":"yellow","a":10,"g":"bl` +
				`ue","b":56,"d":"red","f":["green","red",` +
				`"red"]},["violet","orange",{"e":7,"c":17` +
				`0,"a":"green","b":55,"d":115},"green","b` +
				`lue"]]]},{"e":{"e":{"a":"blue","b":[192,` +
				`"blue",86,93]},"a":"green","d":"yellow",` +
				`"c":186,"h":["violet","orange","orange",` +
				`"violet","red","orange",139,"violet","gr` +
				`een",{"e":"red","a":85,"d":"orange","c":` +
				`"yellow","h":[46,35,"red","green",-11,"b` +
				`lue"],"b":"yellow","g":"yellow","f":"ora` +
				`nge"}],"b":{"e":77,"a":"yellow","d":"gre` +
				`en","c":144,"h":"green","b":{"e":[27,"bl` +
				`ue","yellow",-48,-21,-12,121,"violet"],"` +
				`a":[-23],"d":"blue","j":{"e":"orange","c` +
				`":"blue","a":"green","b":-34,"d":"green"` +
				`},"c":"green","h":"green","b":61,"g":["b` +
				`lue"],"f":19,"i":"violet"},"g":86,"f":"o` +
				`range"},"g":"yellow","f":"orange"},"c":{` +
				`"c":["yellow",82,"red","orange",{"e":34,` +
				`"c":"green","a":"violet","b":182,"d":"or` +
				`ange","f":{"e":-49,"a":184,"d":57,"j":"y` +
				`ellow","c":120,"h":"violet","b":170,"g":` +
				`159,"f":-3,"i":99}},-37,{"e":84,"a":["vi` +
				`olet",154,"violet",123,"violet",148,105,` +
				`"yellow",195],"d":"orange","c":{"a":140}` +
				`,"h":"yellow","b":159,"g":76,"f":186},18` +
				`3],"a":[{"e":"yellow","a":-2,"d":"green"` +
				`,"c":{"c":68,"a":"red","b":"blue"},"h":[` +
				`140,99,-2,"green","orange","orange",-14,` +
				`60,"red","green"],"b":"violet","g":1,"f"` +
				`:["yellow","violet"],"i":142},{"e":{"a":` +
				`76,"b":-17},"c":46,"a":[76,1,79,36,-25,"` +
				`yellow",0],"g":"yellow","b":185,"d":54,"` +
				`f":"green"}],"b":{"e":127,"a":[["violet"` +
				`],115,114,"red","orange",83,-17,-2],"d":` +
				`{"e":198,"c":"red","a":46,"b":77,"d":"gr` +
				`een"},"c":140,"h":"orange","b":89,"g":14` +
				`9,"f":"orange"}},"a":"blue","g":62,"b":"` +
				`orange","d":"violet","f":{"e":-11,"a":[2` +
				`9,"yellow","yellow",187,"orange",{"e":14` +
				`7,"a":197,"d":["green",182,-2,95,-8,110,` +
				`-38],"c":"violet","h":187,"b":90,"g":22,` +
				`"f":"yellow","i":"green"},{"a":["blue","` +
				`red",140],"b":"violet"},"blue",76,59],"d` +
				`":-26,"c":[{"a":"orange","b":179},"red",` +
				`{"e":"violet","c":"orange","a":"blue","g` +
				`":"violet","b":25,"d":149,"f":-27},{"e":` +
				`"green","a":"yellow","d":"violet","j":{"` +
				`c":67,"a":179,"b":53},"c":145,"h":-4,"b"` +
				`:"blue","g":11,"f":"blue","i":"violet"}]` +
				`,"h":"red","b":"blue","g":{"e":"yellow",` +
				`"c":[93,12,118,-7,125,93,"yellow",182,11` +
				`3,"yellow"],"a":"green","b":"green","d":` +
				`"violet"},"f":"green","i":"green"}},[7]]` +
				`,"d":[[[42,67,{"e":"yellow","a":"orange"` +
				`,"d":{"e":"orange","c":"green","a":-26,"` +
				`g":[68,162,"orange","red"],"b":105,"d":5` +
				`2,"f":"yellow"},"c":{"e":"orange","a":"o` +
				`range","d":120,"c":"blue","h":[172,"oran` +
				`ge",171,-40,139,161,"yellow",197],"b":"y` +
				`ellow","g":[162,"orange",2],"f":"yellow"` +
				`,"i":"orange"},"h":-21,"b":"green","g":{` +
				`"e":33,"a":6,"d":"violet","c":193,"h":89` +
				`,"b":56,"g":146,"f":{"c":-46,"a":"green"` +
				`,"b":161},"i":149},"f":65},[["yellow",31` +
				`,-4],"red",-27,21,{"e":{"e":-25,"a":-1,"` +
				`d":"violet","c":"violet","h":"blue","b":` +
				`"green","g":"violet","f":"orange","i":"y` +
				`ellow"},"c":"orange","a":"orange","g":"o` +
				`range","b":60,"d":{"e":"yellow","c":"gre` +
				`en","a":"yellow","g":"violet","b":134,"d` +
				`":149,"f":"yellow"},"f":"blue"},"violet"` +
				`,[66,"green",25,106,"red","orange"],{"a"` +
				`:{"e":"orange","a":95,"d":"red","c":176,` +
				`"h":179,"b":87,"g":195,"f":71,"i":"viole` +
				`t"},"b":14},55],"blue"],[[62,"violet",58` +
				`,[-42,"orange",{"e":186,"a":"orange","d"` +
				`:"green","j":100,"c":163,"h":-9,"b":"gre` +
				`en","g":"orange","f":-8,"i":"orange"},"r` +
				`ed",12,"orange",-3,162,"green"],"violet"` +
				`,94,"orange",41,58,"violet"],[["green","` +
				`green","red"],"red",128,"blue","yellow",` +
				`"yellow",{"e":"yellow","a":193,"d":["ora` +
				`nge",134],"j":["green",45,195,123,50,61]` +
				`,"c":"violet","h":-39,"b":13,"g":"yellow` +
				`","f":{"e":-33,"c":26,"a":83,"b":122,"d"` +
				`:"orange"},"i":"orange"},"blue",46,"yell` +
				`ow"],"yellow","yellow",30,[158,{"c":"gre` +
				`en","a":68,"b":122},{"e":95,"c":"orange"` +
				`,"a":"green","b":115,"d":180,"f":"yellow` +
				`"},[194,-29,"orange","violet",{"e":173,"` +
				`a":193,"d":"green","j":"blue","c":"blue"` +
				`,"h":"green","b":"yellow","g":"green","f` +
				`":186,"i":"red"},135],"violet"]],"yellow` +
				`",144,["violet","red",{"a":[121],"b":134` +
				`},{"a":-23,"b":"violet"},[7],[101,181,"y` +
				`ellow",{"e":"violet","c":"green","a":56,` +
				`"b":"green","d":175}]],"violet",{"c":"gr` +
				`een","a":{"e":"green","c":-42,"a":-49,"b` +
				`":[["orange","yellow",-17,"orange",-11,-` +
				`41,"red",32],"blue",46,{"a":"yellow","b"` +
				`:"violet"}],"d":"red","f":["violet","vio` +
				`let",-11]},"b":["orange",{"e":62,"c":"vi` +
				`olet","a":-5,"b":39,"d":[122,129,"violet` +
				`","orange"]},"green"]}]],"f":{"e":{"e":[` +
				`{"e":{"a":"violet"},"c":171,"a":{"a":{"e` +
				`":122,"a":"blue","d":164,"j":54,"c":"vio` +
				`let","h":-1,"b":148,"g":"blue","f":190,"` +
				`i":-37}},"b":[{"a":158},"blue",26,{"e":"` +
				`violet","c":"green","a":64,"g":36,"b":"b` +
				`lue","d":"blue","f":"green"},"violet",12` +
				`5,"orange"],"d":"violet"},"yellow",{"c":` +
				`{"e":-46,"c":121,"a":191,"g":"blue","b":` +
				`67,"d":-25,"f":"blue"},"a":"orange","b":` +
				`{"a":108,"b":["red",-48]}},{"e":116,"c":` +
				`"green","a":-34,"g":59,"b":"yellow","d":` +
				`37,"f":{"e":-3,"a":105,"d":"red","c":"re` +
				`d","h":"green","b":27,"g":189,"f":"viole` +
				`t"}},"red",["yellow",152,{"e":"red","a":` +
				`[164,155,"green","yellow"],"d":79,"c":"v` +
				`iolet","h":{"e":"blue","c":53,"a":"orang` +
				`e","b":-43,"d":"violet","f":"orange"},"b` +
				`":32,"g":"green","f":"green","i":23},"re` +
				`d","violet",-37,[["green","violet",131,-` +
				`18,"green"],48,6,"red",83],152,181,75],{` +
				`"e":[116,100,"red",{"a":156},"green"],"c` +
				`":"red","a":170,"g":[28,59,"violet","vio` +
				`let",144],"b":13,"d":{"a":"yellow","b":"` +
				`blue"},"f":{"e":-45,"c":"orange","a":"or` +
				`ange","b":"green","d":["orange","orange"` +
				`,98,-13,"yellow",183,193,"green"],"f":"g` +
				`reen"}},128,["green"]],"a":{"e":{"e":"gr` +
				`een","a":"yellow","d":{"c":-48,"a":["vio` +
				`let",195,77,-13,"green","red"],"b":"blue` +
				`","d":"red"},"c":"violet","h":[-13,"gree` +
				`n"],"b":27,"g":-37,"f":197},"c":68,"a":-` +
				`23,"b":93,"d":"orange","f":{"a":"green",` +
				`"b":-4}},"d":[-48],"c":[["violet","blue"` +
				`,"orange",{"e":42,"c":66,"a":["blue","bl` +
				`ue","orange",144,130,"red","violet","red` +
				`",112],"b":"blue","d":192},"orange",["or` +
				`ange",4,"violet",-47,"orange"],["violet"` +
				`,"yellow",54,{"c":47,"a":"yellow","b":"r` +
				`ed"},{"a":"violet"},151,-30,"orange","re` +
				`d"],-31,"blue"],135,{"e":"red","a":[5,11` +
				`7],"d":119,"j":189,"c":{"e":134,"a":142,` +
				`"d":"orange","j":189,"c":"red","h":"viol` +
				`et","b":"yellow","g":122,"f":{"c":"orang` +
				`e","a":"violet","b":"yellow"},"i":107},"` +
				`h":109,"b":15,"g":[157,135,122,-34,168,"` +
				`green",-19,13,64],"f":{"e":"blue","c":17` +
				`4,"a":187,"g":"violet","b":36,"d":80,"f"` +
				`:188},"i":-10},60,"orange",146,{"e":-8,"` +
				`a":{"a":63},"d":"green","c":67,"h":["blu` +
				`e","yellow","red","blue",["yellow",189,"` +
				`orange","green"],"orange","violet"],"b":` +
				`-29,"g":"orange","f":["violet","green","` +
				`orange"]},{"e":"violet","c":"red","a":-4` +
				`5,"g":[45,"green",-24,-7,{"a":67,"b":"re` +
				`d"},"orange",188,["red","blue","red"],47` +
				`,133],"b":{"e":112,"a":76,"d":"yellow","` +
				`j":93,"c":"orange","h":137,"b":154,"g":[` +
				`"orange","red","yellow",-26,"orange"],"f` +
				`":-3,"i":"yellow"},"d":{"e":"violet","c"` +
				`:48,"a":174,"b":"green","d":189},"f":[69` +
				`,"red",{"a":"green","b":"blue"},"red",{"` +
				`a":85}]},{"e":"violet","c":{"e":"violet"` +
				`,"c":"red","a":-22,"b":3,"d":{"e":185,"c` +
				`":52,"a":"yellow","b":165,"d":"blue","f"` +
				`:"yellow"}},"a":45,"b":"green","d":"blue` +
				`"}],"h":"orange","b":{"e":"red","c":68,"` +
				`a":171,"g":-30,"b":"orange","d":"violet"` +
				`,"f":{"c":22,"a":["green"],"b":148}},"g"` +
				`:{"e":77,"c":28,"a":["green","green",{"a` +
				`":"violet","b":115},{"e":59,"a":198,"d":` +
				`"violet","c":"yellow","h":-38,"b":"viole` +
				`t","g":69,"f":"violet","i":-12},35],"b":` +
				`"blue","d":{"e":[-20,"orange","green",11` +
				`6,"yellow",-21],"c":"red","a":"green","b` +
				`":160,"d":["violet",-49,"yellow",[46,"bl` +
				`ue",196]]}},"f":[162,{"c":[-30,"violet"]` +
				`,"a":"blue","b":"blue","d":83},"blue"],"` +
				`i":"green"},"a":["blue",{"e":{"e":131,"c` +
				`":"yellow","a":"blue","g":{"e":"yellow",` +
				`"c":70,"a":{"e":53,"c":100,"a":162,"b":-` +
				`7,"d":-14,"f":"orange"},"b":"yellow","d"` +
				`:"green","f":149},"b":[184,"green","blue` +
				`",-38,193,75,156],"d":96,"f":{"e":177,"c` +
				`":0,"a":{"c":"red","a":-23,"b":"yellow",` +
				`"d":-37},"g":149,"b":"blue","d":61,"f":7` +
				`8}},"c":-21,"a":[{"e":{"a":"yellow"},"c"` +
				`:71,"a":{"e":-46,"a":"yellow","d":"green` +
				`","c":78,"h":59,"b":"green","g":189,"f":` +
				`"red","i":"red"},"b":"yellow","d":{"e":"` +
				`violet","c":"violet","a":"orange","b":18` +
				`9,"d":21,"f":71},"f":15},"green",150,"ye` +
				`llow"],"b":101,"d":"blue","f":"blue"},"g` +
				`reen",[[-10,{"c":61,"a":99,"b":"green"}]` +
				`,"violet",["blue",-41],54,[194,146,"gree` +
				`n",90,"violet","violet",-5,"red"],{"e":1` +
				`14,"a":"green","d":{"e":"blue","a":"blue` +
				`","d":["yellow",10,"blue",86,44,"violet"` +
				`,"red","red"],"c":94,"h":{"e":"red","c":` +
				`"blue","a":89,"g":197,"b":148,"d":91,"f"` +
				`:52},"b":"orange","g":"red","f":"green"}` +
				`,"c":[96,127,29],"h":"blue","b":"green",` +
				`"g":"yellow","f":{"e":"green","a":"blue"` +
				`,"d":"orange","j":40,"c":"red","h":92,"b` +
				`":-49,"g":-34,"f":8,"i":125},"i":-30},{"` +
				`a":154,"b":100}]],"d":{"c":{"a":"violet"` +
				`,"b":"violet"},"a":"yellow","b":{"e":[83` +
				`,"red",["green",{"a":"violet"},126,105,8` +
				`6],[["blue",-40,148,"blue"],168,"red","g` +
				`reen",["green",41,93,"red",-19],32,"viol` +
				`et","violet","blue","yellow"],{"a":"yell` +
				`ow","b":"green"},[85,164],17,60,66,{"c":` +
				`"yellow","a":"yellow","b":-10,"d":"yello` +
				`w"}],"a":[109,28,52,[118,"violet"],{"a":` +
				`"blue"},10,163,"green",-21],"d":"green",` +
				`"c":"green","h":["violet","violet"],"b":` +
				`"blue","g":{"e":-22,"c":["yellow",119,12` +
				`7],"a":"green","g":1,"b":[{"a":100,"b":"` +
				`yellow"},63,41,168,152,"yellow",198,-14,` +
				`30,103],"d":"violet","f":["blue",135,"ye` +
				`llow","green","yellow",["yellow",35,91,1` +
				`82]]},"f":"yellow","i":{"c":"green","a":` +
				`[{"a":"yellow"},"green","orange",50,117,` +
				`94,"red",89],"b":"red"}}},"j":{"e":{"e":` +
				`{"a":70},"c":["blue","violet"],"a":113,"` +
				`b":{"e":"yellow","c":-9,"a":135,"b":"yel` +
				`low","d":81},"d":{"e":189,"a":-19,"d":-1` +
				`4,"j":{"e":"green","a":1,"d":"violet","c` +
				`":"red","h":106,"b":"green","g":-19,"f":` +
				`192,"i":"green"},"c":195,"h":"yellow","b` +
				`":"orange","g":45,"f":"green","i":"viole` +
				`t"},"f":75},"a":"violet","d":"blue","c":` +
				`"red","h":113,"b":[14,{"e":82,"a":{"e":5` +
				`9,"a":182,"d":"yellow","c":"blue","h":"y` +
				`ellow","b":"green","g":93,"f":"yellow","` +
				`i":147},"d":58,"c":"violet","h":{"a":{"e` +
				`":27,"a":"orange","d":181,"c":142,"h":19` +
				`5,"b":"yellow","g":44,"f":"yellow"}},"b"` +
				`:-14,"g":"red","f":{"a":154},"i":90},"or` +
				`ange",{"a":"blue","b":-11},["yellow",[38` +
				`,-34,"orange",57],"orange",98,"violet","` +
				`violet",{"a":134},[24,"blue","blue",172,` +
				`114],10],[{"a":["yellow","green"]},166,"` +
				`blue",["blue",{"e":"blue","c":-6,"a":"gr` +
				`een","b":"green","d":"green","f":50},-44` +
				`,{"e":"violet","a":"yellow","d":-27,"j":` +
				`33,"c":"orange","h":146,"b":"green","g":` +
				`30,"f":"violet","i":"violet"},"blue","re` +
				`d","violet",6]],113],"g":{"e":{"e":"red"` +
				`,"a":125,"d":183,"j":74,"c":["green","ye` +
				`llow"],"h":12,"b":["green",-38,"yellow"]` +
				`,"g":71,"f":"blue","i":"blue"},"a":["gre` +
				`en",86,28,[55,127,"blue",172,"green",83,` +
				`"green","violet",{"a":53},"blue"],{"a":1` +
				`81,"b":57},[28,"violet"],"red",{"a":138}` +
				`,198],"d":"yellow","c":[{"e":"violet","c` +
				`":65,"a":"green","b":{"e":"orange","a":6` +
				`8,"d":"orange","c":"green","h":161,"b":"` +
				`green","g":"green","f":13},"d":-38},"blu` +
				`e",{"e":73,"a":40,"d":"green","c":"red",` +
				`"h":"yellow","b":117,"g":178,"f":170},"y` +
				`ellow",188],"h":{"a":-21},"b":{"c":86,"a` +
				`":["green","green",170,"blue",-40,"yello` +
				`w","red",10],"b":[["green","yellow",2,"v` +
				`iolet",-30,"yellow","red",-40,"orange","` +
				`violet"],"red",58,165,"red"],"d":13},"g"` +
				`:{"a":63},"f":"violet"},"f":["green"],"i` +
				`":[["yellow","yellow","red",{"e":-5,"c":` +
				`124,"a":"yellow","g":"violet","b":"orang` +
				`e","d":"violet","f":149},-16,9],"violet"` +
				`,-43,{"c":9,"a":52,"b":179,"d":74},"viol` +
				`et",[{"e":-30,"a":82,"d":"green","j":49,` +
				`"c":{"e":"green","a":"violet","d":"yello` +
				`w","j":"blue","c":37,"h":"violet","b":48` +
				`,"g":12,"f":"green","i":"red"},"h":58,"b` +
				`":["yellow","blue","blue",29,"orange","g` +
				`reen","green",0],"g":128,"f":"orange","i` +
				`":91},[181,"red","green","violet","red",` +
				`"green","orange",{"e":"green","a":"yello` +
				`w","d":"green","c":24,"h":194,"b":128,"g` +
				`":"red","f":162,"i":"violet"},191],"red"` +
				`,"violet","yellow","red"],141,123]},"c":` +
				`{"a":95},"h":{"a":138},"b":118,"g":"gree` +
				`n","f":0,"i":"violet"}}`,
			Result1: "111754",
			Result2: "65402"},
	}
	for _, tt := range testCases {
		tt.Test(Day12, assert)
	}
}
