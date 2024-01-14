package mocks

import (
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/AkshachRd/leards-backend-go/services"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

var db *gorm.DB

type MockUserData struct {
	name     string
	email    string
	password string
}

func getMockedUsers() []MockUserData {
	return []MockUserData{
		{
			name:     "Иван Петров",
			email:    "ivan@mail.com",
			password: "12345Q",
		},
		{
			name:     "Мария Иванова",
			email:    "maria@mail.com",
			password: "12345Q",
		},
		{
			name:     "Александр Смирнов",
			email:    "alex@mail.com",
			password: "12345Q",
		},
		{
			name:     "Екатерина Николаева",
			email:    "ekaterina@mail.com",
			password: "12345Q",
		},
		{
			name:     "Павел Козлов",
			email:    "pavel@mail.com",
			password: "12345Q",
		},
		{
			name:     "Ольга Сергеева",
			email:    "olga@mail.com",
			password: "12345Q",
		},
		{
			name:     "Дмитрий Александров",
			email:    "dmitriy@mail.com",
			password: "12345Q",
		},
		{
			name:     "Наталья Кузнецова",
			email:    "natalia@mail.com",
			password: "12345Q",
		},
		{
			name:     "Артем Владимиров",
			email:    "artem@mail.com",
			password: "12345Q",
		},
		{
			name:     "Елена Игнатьева",
			email:    "elena@mail.com",
			password: "12345Q",
		},
		{
			name:     "Сергей Павлов",
			email:    "sergey@mail.com",
			password: "12345Q",
		},
		{
			name:     "Анна Сидорова",
			email:    "anna@mail.com",
			password: "12345Q",
		},
		{
			name:     "Владислава Лебедева",
			email:    "vladislava@mail.com",
			password: "12345Q",
		},
		{
			name:     "Максим Федоров",
			email:    "max@mail.com",
			password: "12345Q",
		},
		{
			name:     "Юлия Антонова",
			email:    "yulia@mail.com",
			password: "12345Q",
		},
		{
			name:     "Денис Морозов",
			email:    "denis@mail.com",
			password: "12345Q",
		},
		{
			name:     "Евгения Краснова",
			email:    "evgenia@mail.com",
			password: "12345Q",
		},
		{
			name:     "Игорь Степанов",
			email:    "igor@mail.com",
			password: "12345Q",
		},
		{
			name:     "Анастасия Васнецова",
			email:    "anastasia@mail.com",
			password: "12345Q",
		},
		{
			name:     "Сергей Захаров",
			email:    "sergei@mail.com",
			password: "12345Q",
		},
	}
}

func MockData(dbToMock *gorm.DB) {
	db = dbToMock
	for _, user := range getMockedUsers() {
		err := MockUserWithDefaultContent(user.name, user.email, user.password)
		if err != nil {
			log.Fatalf("Error mocking up data in database: %v", err)
		}
	}
}

func MockUserWithDefaultContent(name string, email string, password string) error {
	rand.Seed(time.Now().UnixNano())
	_, err := models.FetchUserByEmail(email)
	if err == nil {
		return nil
	}

	user, err := models.CreateUser(name, email, password)
	if err != nil {
		return err
	}

	min := 3
	max := 17
	deckCount := rand.Intn(max-min+1) + min

	for i := 0; i < deckCount; i++ {
		err = createMockDeck(user)

		if err != nil {
			return err
		}
	}

	return nil
}

func createMockDeck(user *models.User) error {
	deck, err := models.CreateDeck(generateRandomDeckName(), models.AccessTypePublic, user.RootFolderID, user.ID)

	if err != nil {
		return err
	}

	fillDeck(deck)

	err = db.Save(&deck).Error

	if err != nil {
		return err
	}

	err = addTagsToMockedDeck(deck)

	if err != nil {
		return err
	}

	return nil
}

func fillDeck(deck *models.Deck) {
	min := 8
	max := 50
	cardsCount := rand.Intn(max-min+1) + min

	for i := 0; i < cardsCount; i++ {
		frontSide := generateRandomWord()
		backSide := translateToRussian(frontSide)

		card := models.Card{
			DeckID:    deck.ID,
			FrontSide: frontSide,
			BackSide:  backSide,
		}

		deck.Cards = append(deck.Cards, card)
	}
}

func generateRandomDeckName() string {
	words := []string{
		"Вдохновение и творчество", "Энергия и гармония", "Исследование мира", "Моменты прозрения", "Сила в мелочах",
		"Цвета жизни", "Разнообразие перспектив", "Путешествие мысли", "Волшебство в обыденности", "Слова в картинах",
		"Ароматы времени", "Мелодии эмоций", "Символы судьбы", "Отражение души", "Ритмы времени", "Тайны природы",
		"Случайности смысла", "Ключи к пониманию", "Интеллектуальные стимулы", "Эмоциональный лабиринт", "Звуки молчания",
		"Точки зрения", "Поэзия звуков", "Очарование взгляда", "Лингвистический пазл", "Симфония мыслей", "Жизнь в цветах",
		"Абстракции реальности", "Калейдоскоп идей", "Магия слова", "Гармония элементов", "Экспрессия времени", "Игра теней",
		"Мозаика моментов", "Картины сознания", "Реальность фантазии", "Этюды смысла", "Тайные коды", "Метафоры ушедшего",
		"Следы ветра", "Моментальная проза", "Орнаменты разума", "Истории облаков", "Точка равновесия", "Палитра воспоминаний",
		"Потоки сознания", "Эпохи взгляда", "Орден хаоса", "Разговор с молчанием", "Секреты времени", "Ассоциации памяти",
		"Контрасты восприятия", "Эклектика мысли", "Хроники вдохновения", "Фрагменты мира", "Плавание в идеях", "Созвучие сущности",
		"Алхимия взгляда", "Проект мозаики", "Моменты замедления", "Гиперреальность души", "Отражение мгновений",
		"Сюрреализм бытия", "Символы мудрости", "Реальность гармонии", "Лабиринты воспоминаний", "Закатные размышления",
		"Архитектура снов", "Поэтика времени", "Эффект бабочки", "Игра в противоречия", "Визуальные гармонии",
		"Искусство открытий", "Магия смысла", "Созвучие времени", "Реальность утопий", "Ассоциации сновидений",
		"Орнаменты мысли", "Эмоциональная симфония", "Искусство понимания", "Звуки вдохновения", "Эклектика переживаний",
		"Поэзия взаимодействия", "Мозаика мудрости", "Абстрактные ландшафты", "Словесные пейзажи", "Метафоры времени",
		"Этюды разума", "Моменты чуда", "Игра смысла", "Теория цвета мысли", "Созвучие чувств", "Реальность фантазии",
		"Мозаика звуков", "Симфония разума", "Словесная галерея", "Эффект момента", "Орнаменты восприятия", "Абстракции чувств",
		"Гармония вдохновения", "Игра света", "Ритмы разума", "Эмоциональные аккорды", "Палитра мысли", "Картины времени",
		"Словесные ароматы", "Мозаика воспоминаний", "Симфония чувств", "Разговор с прошлым", "Абстракции эмоций", "Исследование себя",
		"Танец идей", "Экспрессия взгляда", "Мозаика переживаний", "Гармония вопросов", "Эмоциональные пейзажи", "Абстрактные моменты",
		"Словесные зарисовки", "Исследование вдохновения", "Магия момента", "Реальность чувств", "Симфония воспоминаний", "Эффект слова",
		"Орнаменты вдохновения", "Абстракции восприятия", "Игра смысла", "Теория цвета мысли", "Созвучие чувств", "Реальность фантазии",
		"Мозаика звуков", "Симфония разума", "Словесная галерея", "Эффект момента", "Орнаменты восприятия",
		"Абстракции чувств", "Гармония вдохновения", "Игра света", "Ритмы разума", "Эмоциональные аккорды", "Палитра мысли",
		"Картины времени", "Словесные ароматы", "Мозаика воспоминаний", "Симфония чувств", "Разговор с прошлым",
		"Абстракции эмоций", "Исследование себя", "Танец идей", "Экспрессия взгляда", "Мозаика переживаний", "Гармония вопросов",
		"Эмоциональные пейзажи", "Абстрактные моменты", "Словесные зарисовки", "Исследование вдохновения", "Магия момента",
		"Реальность чувств", "Симфония воспоминаний", "Эффект слова", "Орнаменты вдохновения", "Абстракции восприятия",
	}
	return words[rand.Intn(len(words))]
}

func generateRandomWord() string {
	words := []string{
		"Apple", "Banana", "Orange", "Grape", "Pineapple", "Lemon", "Watermelon", "Strawberry", "Cherry", "Blueberry",
		"Mango", "Peach", "Plum", "Raspberry", "Blackberry", "Cucumber", "Carrot", "Tomato", "Broccoli", "Spinach",
		"Pumpkin", "Potato", "Onion", "Garlic", "Cabbage", "Cauliflower", "Pepper", "Lettuce", "Cucumber", "Radish",
		"Avocado", "Eggplant", "Zucchini", "Kiwi", "Pomegranate", "Coconut", "Pear", "Grapefruit", "Cantaloupe", "Fig",
		"Quince", "Apricot", "Nectarine", "Clementine", "Lime", "Cranberry", "Gooseberry", "Honeydew", "Persimmon", "Date",
		"Papaya", "Guava", "Lychee", "Rambutan", "Passionfruit", "Dragonfruit", "Kale", "Artichoke", "Asparagus", "Celery",
		"Radish", "Beetroot", "Turnip", "Parsnip", "Squash", "Cranberry", "Blueberry", "Raspberry", "Strawberry", "Blackberry",
		"Cherry", "Ginger", "Turmeric", "Cinnamon", "Vanilla", "Nutmeg", "Coriander", "Basil", "Cilantro", "Mint", "Parsley",
		"Thyme", "Rosemary", "Sage", "Oregano", "Dill", "Chives", "Tarragon", "Cumin", "Cardamom", "Cloves", "Mustard", "BayLeaf",
		"Allspice", "Fennel", "Peppermint", "Lavender", "Chamomile", "Eucalyptus", "Jasmine", "Lemongrass", "Patchouli", "Sandalwood",
	}
	return words[rand.Intn(len(words))]
}

func translateToRussian(word string) string {
	translation := map[string]string{
		"Apple":        "Яблоко",
		"Banana":       "Банан",
		"Orange":       "Апельсин",
		"Grape":        "Виноград",
		"Pineapple":    "Ананас",
		"Lemon":        "Лимон",
		"Watermelon":   "Арбуз",
		"Strawberry":   "Клубника",
		"Cherry":       "Вишня",
		"Blueberry":    "Голубика",
		"Mango":        "Манго",
		"Peach":        "Персик",
		"Plum":         "Слива",
		"Raspberry":    "Малина",
		"Blackberry":   "Ежевика",
		"Cucumber":     "Огурец",
		"Carrot":       "Морковь",
		"Tomato":       "Помидор",
		"Broccoli":     "Брокколи",
		"Spinach":      "Шпинат",
		"Pumpkin":      "Тыква",
		"Potato":       "Картошка",
		"Onion":        "Лук",
		"Garlic":       "Чеснок",
		"Cabbage":      "Капуста",
		"Cauliflower":  "Цветная капуста",
		"Pepper":       "Перец",
		"Lettuce":      "Салат",
		"Radish":       "Редис",
		"Avocado":      "Авокадо",
		"Eggplant":     "Баклажан",
		"Zucchini":     "Кабачок",
		"Kiwi":         "Киви",
		"Pomegranate":  "Гранат",
		"Coconut":      "Кокос",
		"Pear":         "Груша",
		"Grapefruit":   "Грейпфрут",
		"Cantaloupe":   "Канталупа",
		"Fig":          "Инжир",
		"Quince":       "Айва",
		"Apricot":      "Абрикос",
		"Nectarine":    "Нектарин",
		"Clementine":   "Клементин",
		"Lime":         "Лайм",
		"Cranberry":    "Клюква",
		"Gooseberry":   "Крыжовник",
		"Honeydew":     "Дыня",
		"Persimmon":    "Хурма",
		"Date":         "Финик",
		"Papaya":       "Папайя",
		"Guava":        "Гуава",
		"Lychee":       "Личи",
		"Rambutan":     "Рамбутан",
		"Passionfruit": "Маракуйя",
		"Dragonfruit":  "Питахайя",
		"Kale":         "Кейл",
		"Artichoke":    "Артишок",
		"Asparagus":    "Спаржа",
		"Celery":       "Сельдерей",
		"Beetroot":     "Свекла",
		"Turnip":       "Репа",
		"Parsnip":      "Пастернак",
		"Squash":       "Кабачок",
		"Ginger":       "Имбирь",
		"Turmeric":     "Куркума",
		"Cinnamon":     "Корица",
		"Vanilla":      "Ваниль",
		"Nutmeg":       "Мускатный орех",
		"Coriander":    "Кориандр",
		"Basil":        "Базилик",
		"Cilantro":     "Кинза",
		"Mint":         "Мята",
		"Parsley":      "Петрушка",
		"Thyme":        "Тимьян",
		"Rosemary":     "Розмарин",
		"Sage":         "Шалфей",
		"Oregano":      "Орегано",
		"Dill":         "Укроп",
		"Chives":       "Лук-порей",
		"Tarragon":     "Тархун",
		"Cumin":        "Тмин",
		"Cardamom":     "Кардамон",
		"Cloves":       "Гвоздика",
		"Mustard":      "Горчица",
		"BayLeaf":      "Лавровый лист",
		"Allspice":     "Инжир",
		"Fennel":       "Фенхель",
		"Peppermint":   "Мелисса",
		"Lavender":     "Лаванда",
		"Chamomile":    "Ромашка",
		"Eucalyptus":   "Эвкалипт",
		"Jasmine":      "Жасмин",
		"Lemongrass":   "Лимонная трава",
		"Patchouli":    "Пачули",
		"Sandalwood":   "Сандал",
	}
	return translation[word]
}

func addTagsToMockedDeck(deck *models.Deck) error {
	tagService, err := services.NewTagService(deck.ID, "deck")

	if err != nil {
		return err
	}

	err = tagService.AddTags(generateRandomTags())

	return err
}

func generateRandomTags() []string {
	words := []string{
		"start", "база", "фрукты", "от_Васяна", "английский",
	}

	// Генерация случайной длины для нового среза
	randomLength := rand.Intn(len(words) + 1)

	// Генерация случайного индекса начала обрезки
	startIndex := rand.Intn(len(words) - randomLength + 1)

	// Обрезка массива
	randomSlice := words[startIndex : startIndex+randomLength]

	return randomSlice
}
