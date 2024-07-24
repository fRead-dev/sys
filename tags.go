package ParserInterface

//###########################################################//

const (
	StatusNil       StatusTag = 0 //Статус не определен
	StatusFull      StatusTag = 1 // Закончено
	StatusFreeze    StatusTag = 2 // Заморожено
	StatusInProcess StatusTag = 3 // В процессе
)

//###########################################################//

const (
	RatingNil       RatingTag = 0   // Rating not determined.
	RatingG         RatingTag = 1   // General audiences. Nudity, sexual scenes, and drug use are absent; violence is minimal; expressions may be used that go beyond polite conversation, but only those commonly found in everyday speech.
	RatingPG        RatingTag = 2   // Parental guidance suggested. Explicit sexual scenes and drug use are absent; nudity, if present, is only very limited; mild profanity may be used and scenes of violence are presented, but only in very moderate amounts.
	RatingPG13      RatingTag = 3   // Parents strongly cautioned. May contain moderate or harsh violence; scenes of nudity may be present; there might be situations with a sexual context; some drug use scenes may occur; occasional use of coarse profanity is possible.
	RatingR         RatingTag = 50  // Restricted. An R rating may be assigned due to frequent use of obscene language, prolonged scenes of violence, sexual activity, or drug use.
	RatingNC17      RatingTag = 100 // No One 17 & Under Admitted. The film may contain explicit sexual scenes, a large amount of obscene and sexual language, or scenes of excessive violence. An NC-17 rating, however, does not imply that the film is obscene or pornographic in either the everyday or legal sense of these words.
	RatingUnlimited RatingTag = 200 // No restrictions, any deviations are possible.
)

//###########################################################//

const (
	FocusNil      FocusTag = 0   // Direction is not defined.
	FocusJen      FocusTag = 1   // Gen — sexual relationships are not described or mentioned, nor do they play a decisive role.
	FocusGet      FocusTag = 2   // Het — romantic and/or sexual relationships between a man and a woman.
	FocusSlash    FocusTag = 3   // Slash — romantic and/or sexual relationships between men.
	FocusFemSlash FocusTag = 4   // Femslash — romantic and/or sexual relationships between women.
	FocusOthers   FocusTag = 200 // Other. Does not fit any of the definitions.
)

//###########################################################//
