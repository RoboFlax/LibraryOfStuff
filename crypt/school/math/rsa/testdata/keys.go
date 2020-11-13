package testdata

import (
	"stuff/rsa"
	"stuff/util"
)

var Robin = &rsa.PrivateKey{
	D: util.IntFromString("5389871559774220240180799440912368590479904556479649312508307102961132748499719514597816385443988115934009997012491913208946750652635968366264048278196538801554579038635485773638859928703523799774128028741952782495686359476777987563502601282333845011295430149317499799185105902599400994767507007378169521595"),
	PublicKey: &rsa.PublicKey{
		N: util.IntFromString("132180183489701115413957700574755705909388135551762828378179912286903969784635978572279782785888279986000721355306349300124170313624167795648856422060534188893386976078895735590174953078802461226769494571572703920943447306678076763957901277498576330826854198349332314323981571977058310888548289588163305532271"),
		E: util.IntFromString("515"),
	},
}
var Justus = &rsa.PublicKey{
	N: util.IntFromString("827210260366768758085244340579330206873250671138523829689703441259089637714036490277572255947620218493"),
	E: util.IntFromString("205"),
}
var Yang = &rsa.PublicKey{
	N: util.IntFromString("7739611342427769114707593244041892685173202718515229414521761553568616071354836942198673517320233123191"),
	E: util.IntFromString("203"),
}
var Maik = &rsa.PublicKey{
	N: util.IntFromString("404900137664651066874043"),
	E: util.IntFromString("431"),
}
