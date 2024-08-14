# Go nullable generic types without pointers

Пакет null реализует обобщённый nullable-тип null.Null\[any].

Тип null.Null основан на универсальном типе sql.Null, но имеет закрытую реализацию.
Он ориентирован на API и конфигурацию и обеспечивает сериализацию/десериализацию JSON и YAML (реализация YAML использует пакет gopkg.in/yaml.v3).

Проверка равенства реализована не методом, а функцией Equal, поддерживающей подмножество типов null.Null\[comparable].
