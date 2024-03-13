# Умова

Реалізувати програму-імітатор простого менеджера пам’яті. Користувач
вводить кількість комірок пам’яті та максимальну кількість комірок для
виводу в одному рядку консолі з клавіатури. Після цього робота програми
стає циклічною. Для виконання користувач може вводити 5 команд:
allocate &lt;num_cells&gt; – виділити блок пам’яті розміром &lt;num_cells&gt;
комірок (команда виводить в консоль номер першої комірки виділеного
блоку – &lt;block_id&gt;), free &lt;block_id&gt; – звільнити блок, print –
вивести структуру блоків пам’яті у консоль, exit – завершити роботу
програми, help – вивести коротку інформацію про команди. Намагатись
оптимізувати продуктивність алгоритму виділення блоків.

Приклад роботи з програмою:

Please set memory size and max output width:

&gt;30 10 \
Type &#39;help&#39; for additional info.\
&gt;help\
Available commands:\
help - show this help\
exit - exit this program\
print - print memory blocks map\
allocate &lt;num&gt; - allocate &lt;num&gt; cells. Returns\
block first cell number\
free &lt;num&gt; - free block with first cell number\
&lt;num&gt;\
&gt;print\
|                 |\
|                 |\
|                 |\
&gt;allocate 5\
0\
&gt;print\
|0xxxxxxxx|       |\
|                 |\
|                 |\
&gt;allocate 3\
5\
&gt;print\
|0xxxxxxxx|5xxxx| |\
|                 |\
|                 |\
&gt;allocate 4\
8\

2\

&gt;allocate 10\
12\
&gt;print\
|0xxxxxxxx|5xxxx|8xx|\
|xxx|12xxxxxxxxxxxxx|\
|xxx|               |\
&gt;free 5\
&gt;print\
|0xxxxxxxx|     |8xx|\
|xxx|12xxxxxxxxxxxxx|\
|xxx|               |\
&gt;free 8\
&gt;print\
|0xxxxxxxx|         |\
|   |12xxxxxxxxxxxxx|\
|xxx|               |\
&gt;allocate 6\
5\
&gt;print\
|0xxxxxxxx|5xxxxxxxx|\
|x| |12xxxxxxxxxxxxx|\
|xxx|               |\
&gt;allocate 3\
22\
&gt;print\
|0xxxxxxxx|5xxxxxxxx|\
|x| |12xxxxxxxxxxxxx|\
|xxx|22xxx|         |\
&gt;exit\