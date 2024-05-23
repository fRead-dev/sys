package temp

//###########################################################//

type StatusTag byte

const (
	StatusNil       StatusTag = 0 // Статус не определен
	StatusFull      StatusTag = 4 // Закончено
	StatusFreeze    StatusTag = 6 // Заморожено
	StatusInProcess StatusTag = 8 // В процессе
)

//###########################################################//

type RatingTag byte

const (
	RatingNil       RatingTag = 0  // Рейтинг не определен
	RatingG         RatingTag = 1  // General audiences. Обнажение, сексуальные сцены и сцены приёма наркотиков отсутствуют; насилие минимально; могут употребляться выражения, выходящие за пределы вежливой беседы, но только те, которые постоянно встречаются в повседневной речи.
	RatingPG        RatingTag = 3  // Parental guidance suggested. Явные сексуальные сцены и сцены употребления наркотиков отсутствуют; нагота, если присутствует, только в очень ограниченной степени, могут быть использованы лёгкие ругательства и представлены сцены насилия, но только в очень умеренных количествах.
	RatingPG13      RatingTag = 5  // Parents strongly cautioned. Может присутствовать умеренное или грубое насилие; могут присутствовать сцены с наготой; возможны ситуации с сексуальным контекстом; могут присутствовать некоторые сцены употребления наркотиков; можно услышать единичные употребления грубых ругательств.
	RatingR         RatingTag = 7  // Restricted. Рейтинг R может быть назначен из-за частого употребления непристойной лексики, продолжительных сцен насилия, полового акта или употребления наркотиков.
	RatingNC17      RatingTag = 9  // No One 17 & Under Admitted. Фильм может содержать явные сексуальные сцены, большое количество непристойной и сексуальной лексики, или сцен чрезмерного насилия. Рейтинг NC-17, однако, ещё не означает, что данный фильм является непристойным или порнографическим, как в повседневном, так и в юридическом смысле этих слов.
	RatingUnlimited RatingTag = 11 // Без ограничений, возможна любая жесть
)

//###########################################################//

type FocusTag byte

const (
	FocusNil      FocusTag = 0  // Направленность не определена
	FocusJen      FocusTag = 20 // Джен — не описываются и не упоминаются сексуальные взаимоотношения или они не играют решающей роли.
	FocusGet      FocusTag = 22 // Гет — романтические и/или сексуальные отношения между мужчиной и женщиной.
	FocusSlash    FocusTag = 24 // Слэш — романтические и/или сексуальные взаимоотношения между мужчинами.
	FocusFemSlash FocusTag = 26 // Фемслэш — романтические и/или сексуальные взаимоотношения между женщинами.
	FocusOthers   FocusTag = 28 // Иное. Не подходит ни под одно из определений
)

//###########################################################//
