Wzorce kreacyjne
Spring Boot (Kotlin)

Proszę stworzyć prosty serwis do autoryzacji, który zasymuluje
autoryzację użytkownika za pomocą przesłanej nazwy użytkownika oraz
hasła. Serwis powinien zostać wstrzyknięty do kontrolera za pomocą
anotacji @Autowired. Aplikacja ma oczywiście zawierać jeden kontroler
i powinna zostać napisana w języku Kotlin. Oparta powinna zostać na
frameworku Spring Boot, podobnie jak na zajęciach. Serwis do
autoryzacji powinien być singletonem.

3.0 Należy stworzyć jeden kontroler wraz z danymi wyświetlanymi z
listy na endpoint’cie w formacie JSON - Kotlin + Spring Boot
3.5 Należy stworzyć klasę do autoryzacji (mock) jako Singleton w
formie eager
4.0 Należy obsłużyć dane autoryzacji przekazywane przez użytkownika
4.5 Należy wstrzyknąć singleton do głównej klasy via @Autowired
5.0 Obok wersji Eager do wyboru powinna być wersja Singletona w wersji
lazy
