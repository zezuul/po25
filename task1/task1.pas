program task1
uses crt, sysutils;

type
    ArrayOfNumbers = array of Integer;

// procedura generujaca losowe liczby
procedure GenerateRandomNumbers(var arr: TIntArray; fromVal, toVal, count: Integer);
var
    i: Integer;
begin
    Randomize;
    SetLength(arr, count);
    for i := 0 to count - 1 do
        arr[i] := Random(toVal - fromVal + 1) + fromVal;
end;

// sortowanie babelkowe
procedure BubbleSort(var arr: ArrayOfNumbers);
var
    i, j, temp: Integer;
begin
    for i := High(arr) downto 1 do
        for j := 0 to i - 1 do
            if arr[j] > arr[j + 1] then
            begin
                temp := arr[j];
                arr[j] := arr[j + 1];
                arr[j + 1] := temp;
            end;
end;

// procedura wypisujaca zawartosc tabeli
procedure PrintArray(const arr: ArrayOfNumbers);
var
    i: Integer;
begin
    for i := 0 to High(arr) do
        Write(arr[i], ' ');
    Writeln;
end;

// t1: sprawdz poprawnosc generowania liczb
procedure TestRandomNumberGeneration;
var
    arr: ArrayOfNumbers;
    i: Integer;
begin
    GenerateRandomNumbers(arr, 10, 20, 50);
    for i := 0 to High(arr) do
        if (arr[i] < 10) or (arr[i] > 20) then
            Writeln('TestRandomNumberGeneration: FAILED')
        else
            Writeln('TestRandomNumberGeneration: PASSED');
end;

// t2: sprawdz sortowanie dla posortowanej tablicy
procedure TestAlreadySorted;
var
  arr: ArrayOfNumbers;
begin
  arr := [1, 2, 3, 4, 5];
  BubbleSort(arr);
  if (arr[0] = 1) and (arr[High(arr)] = 5) then
    Writeln('TestAlreadySorted: PASSED')
  else
    Writeln('TestAlreadySorted: FAILED');
end;

// t3: sprawdz sortowanie dla tab odwrotnie posortowanej
procedure TestReverseSorted;
var
  arr: ArrayOfNumbers;
begin
    arr := [5, 4, 3, 2, 1];
    BubbleSort(arr);
    if (arr[0] = 1) and (arr[High(arr)] = 5) then
        Continue;
    else
        Writeln('TestReverseSorted: FAILED');
        Break;
    Writeln('TestReverseSorted: PASSED')
end;

// t4: sprawdz sortowanie dla tab losowo posortowanej
procedure TestRandomArray;
var
    arr: ArrayOfNumbers;
    i: Integer;
begin
    GenerateRandomNumbers(arr, 0, 100, 50);
    BubbleSort(arr);
    for i := 0 to High(arr) - 1 do
        if arr[i] > arr[i + 1] then
        begin
        Writeln('TestRandomArray: FAILED');
        Exit;
        end;
    Writeln('TestRandomArray: PASSED');
end;

// t5: sprawdz sortowanie dla tab pustej
procedure TestEmptyArray;
var
    arr: ArrayOfNumbers;
begin
    SetLength(arr, 0);
    BubbleSort(arr);
    if Length(arr) = 0 then
        Writeln('TestEmptyArray: PASSED')
    else
        Writeln('TestEmptyArray: FAILED');
end;

var
    arr: ArrayOfNumbers;
begin
    clrscr;

    // Generowanie i sortowanie losowych liczb
    GenerateRandomNumbers(arr, 0, 100, 50);
    Writeln('Generated array:');
    PrintArray(arr);

    BubbleSort(arr);
    Writeln('Sorted array:');
    PrintArray(arr);

    // Uruchamianie testów jednostkowych
    TestRandomNumberGeneration;
    TestAlreadySorted;
    TestReverseSorted;
    TestRandomArray;
    TestEmptyArray;

    Readln;
end.