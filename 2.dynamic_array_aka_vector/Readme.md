# Знакомство с динамическим массивом.
### Нужно реализовать [динамический массив](./vector.go) и его некоторые методы:

1. **Append(item T, rest ...T)**<br> 
добавляет элемент в конец массива;
2. **grow(item T, rest ...T);**<br>
   если для вставляемых элементов при добавлении нет больше места, 
   то массив растет в 2 раза; 
3. **InsertByIndex(index int, item T, rest ...T)**<br>
   добавляет элемент по произвольному индексу массива;
4. **RemoveByIndex(index int) error**<br>
   удаляет элемент по произвольному индексу;
5. **Set(index int, item T) error**<br>
   изменяет элемент по индексу;
6. **GetItem(index int) (T, error)**<br>
   получает элемент по индексу.

### Решить [несколько задач для массива](./exercises)