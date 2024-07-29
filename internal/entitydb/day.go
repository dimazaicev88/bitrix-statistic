package entitydb

import "time"

type Day struct {
	uuid             string
	dateStat         time.Time
	hits             uint32
	hosts            uint32
	sessions         uint32
	events           uint32
	guests           uint32
	newGuests        uint32
	favorites        uint32
	totalHosts       uint32
	amAverageTime    float64
	am1              uint32
	am13             uint32
	am36             uint32
	am69             uint32
	am912            uint32
	am1215           uint32
	am1518           uint32
	am1821           uint32
	am2124           uint32
	am24             uint32
	ahAverageHits    float64
	ah1              uint32
	ah25             uint32
	ah69             uint32
	ah1013           uint32
	ah1417           uint32
	ah1821           uint32
	ah2225           uint32
	ah2629           uint32
	ah3033           uint32
	ah34             uint32
	hourHost0        uint32
	hourHost1        uint32
	hourHost2        uint32
	hourHost3        uint32
	hourHost4        uint32
	hourHost5        uint32
	hourHost6        uint32
	hourHost7        uint32
	hourHost8        uint32
	hourHost9        uint32
	hourHost10       uint32
	hourHost11       uint32
	hourHost12       uint32
	hourHost13       uint32
	hourHost14       uint32
	hourHost15       uint32
	hourHost16       uint32
	hourHost17       uint32
	hourHost18       uint32
	hourHost19       uint32
	hourHost20       uint32
	hourHost21       uint32
	hourHost22       uint32
	hourHost23       uint32
	hourGuest0       uint32
	hourGuest1       uint32
	hourGuest2       uint32
	hourGuest3       uint32
	hourGuest4       uint32
	hourGuest5       uint32
	hourGuest6       uint32
	hourGuest7       uint32
	hourGuest8       uint32
	hourGuest9       uint32
	hourGuest10      uint32
	hourGuest11      uint32
	hourGuest12      uint32
	hourGuest13      uint32
	hourGuest14      uint32
	hourGuest15      uint32
	hourGuest16      uint32
	hourGuest17      uint32
	hourGuest18      uint32
	hourGuest19      uint32
	hourGuest20      uint32
	hourGuest21      uint32
	hourGuest22      uint32
	hourGuest23      uint32
	hourNewGuest0    uint32
	hourNewGuest1    uint32
	hourNewGuest2    uint32
	hourNewGuest3    uint32
	hourNewGuest4    uint32
	hourNewGuest5    uint32
	hourNewGuest6    uint32
	hourNewGuest7    uint32
	hourNewGuest8    uint32
	hourNewGuest9    uint32
	hourNewGuest10   uint32
	hourNewGuest11   uint32
	hourNewGuest12   uint32
	hourNewGuest13   uint32
	hourNewGuest14   uint32
	hourNewGuest15   uint32
	hourNewGuest16   uint32
	hourNewGuest17   uint32
	hourNewGuest18   uint32
	hourNewGuest19   uint32
	hourNewGuest20   uint32
	hourNewGuest21   uint32
	hourNewGuest22   uint32
	hourNewGuest23   uint32
	hourSession0     uint32
	hourSession1     uint32
	hourSession2     uint32
	hourSession3     uint32
	hourSession4     uint32
	hourSession5     uint32
	hourSession6     uint32
	hourSession7     uint32
	hourSession8     uint32
	hourSession9     uint32
	hourSession10    uint32
	hourSession11    uint32
	hourSession12    uint32
	hourSession13    uint32
	hourSession14    uint32
	hourSession15    uint32
	hourSession16    uint32
	hourSession17    uint32
	hourSession18    uint32
	hourSession19    uint32
	hourSession20    uint32
	hourSession21    uint32
	hourSession22    uint32
	hourSession23    uint32
	hourHit0         uint32
	hourHit1         uint32
	hourHit2         uint32
	hourHit3         uint32
	hourHit4         uint32
	hourHit5         uint32
	hourHit6         uint32
	hourHit7         uint32
	hourHit8         uint32
	hourHit9         uint32
	hourHit10        uint32
	hourHit11        uint32
	hourHit12        uint32
	hourHit13        uint32
	hourHit14        uint32
	hourHit15        uint32
	hourHit16        uint32
	hourHit17        uint32
	hourHit18        uint32
	hourHit19        uint32
	hourHit20        uint32
	hourHit21        uint32
	hourHit22        uint32
	hourHit23        uint32
	hourEvent0       uint32
	hourEvent1       uint32
	hourEvent2       uint32
	hourEvent3       uint32
	hourEvent4       uint32
	hourEvent5       uint32
	hourEvent6       uint32
	hourEvent7       uint32
	hourEvent8       uint32
	hourEvent9       uint32
	hourEvent10      uint32
	hourEvent11      uint32
	hourEvent12      uint32
	hourEvent13      uint32
	hourEvent14      uint32
	hourEvent15      uint32
	hourEvent16      uint32
	hourEvent17      uint32
	hourEvent18      uint32
	hourEvent19      uint32
	hourEvent20      uint32
	hourEvent21      uint32
	hourEvent22      uint32
	hourEvent23      uint32
	hourFavorite0    uint32
	hourFavorite1    uint32
	hourFavorite2    uint32
	hourFavorite3    uint32
	hourFavorite4    uint32
	hourFavorite6    uint32
	hourFavorite7    uint32
	hourFavorite5    uint32
	hourFavorite8    uint32
	hourFavorite9    uint32
	hourFavorite10   uint32
	hourFavorite11   uint32
	hourFavorite12   uint32
	hourFavorite13   uint32
	hourFavorite14   uint32
	hourFavorite15   uint32
	hourFavorite16   uint32
	hourFavorite17   uint32
	hourFavorite18   uint32
	hourFavorite19   uint32
	hourFavorite20   uint32
	hourFavorite21   uint32
	hourFavorite22   uint32
	hourFavorite23   uint32
	weekdayHost0     uint32
	weekdayHost1     uint32
	weekdayHost2     uint32
	weekdayHost3     uint32
	weekdayHost4     uint32
	weekdayHost5     uint32
	weekdayHost6     uint32
	weekdayGuest0    uint32
	weekdayGuest1    uint32
	weekdayGuest2    uint32
	weekdayGuest3    uint32
	weekdayGuest4    uint32
	weekdayGuest5    uint32
	weekdayGuest6    uint32
	weekdayNewGuest0 uint32
	weekdayNewGuest1 uint32
	weekdayNewGuest2 uint32
	weekdayNewGuest3 uint32
	weekdayNewGuest4 uint32
	weekdayNewGuest5 uint32
	weekdayNewGuest6 uint32
	weekdaySession0  uint32
	weekdaySession1  uint32
	weekdaySession2  uint32
	weekdaySession3  uint32
	weekdaySession4  uint32
	weekdaySession5  uint32
	weekdaySession6  uint32
	weekdayHit0      uint32
	weekdayHit1      uint32
	weekdayHit2      uint32
	weekdayHit3      uint32
	weekdayHit4      uint32
	weekdayHit5      uint32
	weekdayHit6      uint32
	weekdayEvent0    uint32
	weekdayEvent1    uint32
	weekdayEvent2    uint32
	weekdayEvent3    uint32
	weekdayEvent4    uint32
	weekdayEvent5    uint32
	weekdayEvent6    uint32
	weekdayFavorite0 uint32
	weekdayFavorite1 uint32
	weekdayFavorite2 uint32
	weekdayFavorite3 uint32
	weekdayFavorite4 uint32
	weekdayFavorite5 uint32
	weekdayFavorite6 uint32
	monthHost1       uint32
	monthHost2       uint32
	monthHost3       uint32
	monthHost4       uint32
	monthHost5       uint32
	monthHost6       uint32
	monthHost7       uint32
	monthHost8       uint32
	monthHost9       uint32
	monthHost10      uint32
	monthHost11      uint32
	monthHost12      uint32
	monthGuest1      uint32
	monthGuest2      uint32
	monthGuest3      uint32
	monthGuest4      uint32
	monthGuest5      uint32
	monthGuest6      uint32
	monthGuest7      uint32
	monthGuest8      uint32
	monthGuest9      uint32
	monthGuest10     uint32
	monthGuest11     uint32
	monthGuest12     uint32
	monthNewGuest1   uint32
	monthNewGuest2   uint32
	monthNewGuest3   uint32
	monthNewGuest4   uint32
	monthNewGuest5   uint32
	monthNewGuest6   uint32
	monthNewGuest7   uint32
	monthNewGuest8   uint32
	monthNewGuest9   uint32
	monthNewGuest10  uint32
	monthNewGuest11  uint32
	monthNewGuest12  uint32
	monthSession1    uint32
	monthSession2    uint32
	monthSession3    uint32
	monthSession4    uint32
	monthSession5    uint32
	monthSession6    uint32
	monthSession7    uint32
	monthSession8    uint32
	monthSession9    uint32
	monthSession10   uint32
	monthSession11   uint32
	monthSession12   uint32
	monthHit1        uint32
	monthHit2        uint32
	monthHit3        uint32
	monthHit4        uint32
	monthHit5        uint32
	monthHit6        uint32
	monthHit7        uint32
	monthHit8        uint32
	monthHit9        uint32
	monthHit10       uint32
	monthHit11       uint32
	monthHit12       uint32
	monthEvent1      uint32
	monthEvent2      uint32
	monthEvent3      uint32
	monthEvent4      uint32
	monthEvent5      uint32
	monthEvent6      uint32
	monthEvent7      uint32
	monthEvent8      uint32
	monthEvent9      uint32
	monthEvent10     uint32
	monthEvent11     uint32
	monthEvent12     uint32
	monthFavorite1   uint32
	monthFavorite2   uint32
	monthFavorite3   uint32
	monthFavorite4   uint32
	monthFavorite5   uint32
	monthFavorite6   uint32
	monthFavorite7   uint32
	monthFavorite8   uint32
	monthFavorite9   uint32
	monthFavorite10  uint32
	monthFavorite11  uint32
	monthFavorite12  uint32
}