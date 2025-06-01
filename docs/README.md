<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <title>Методические материалы: Команды терминала Linux</title>
    <style>
        body {
            font-family: 'Arial', sans-serif;
            line-height: 1.6;
            max-width: 900px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f4f4f4;
        }
        .container {
            background-color: white;
            padding: 30px;
            border-radius: 8px;
            box-shadow: 0 4px 6px rgba(0,0,0,0.1);
        }
        h1, h2 {
            color: #333;
            border-bottom: 2px solid #4CAF50;
            padding-bottom: 10px;
        }
        .section {
            margin-bottom: 30px;
            background-color: #f9f9f9;
            padding: 15px;
            border-radius: 5px;
        }
        pre {
            background-color: #f1f1f1;
            border-radius: 4px;
            padding: 15px;
            overflow-x: auto;
        }
        .command-description {
            color: #666;
            font-style: italic;
        }
        table {
            width: 100%;
            border-collapse: collapse;
        }
        table, th, td {
            border: 1px solid #ddd;
            padding: 8px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🖥️ Методические материалы: Команды терминала Linux</h1>

        <div class="section">
            <h2>🎯 Основы навигации в терминале</h2>
            
            <h3>Определение текущего местоположения</h3>
            <pre><code>pwd                    # Показать текущий каталог</code></pre>

            <h3>Перемещение по каталогам</h3>
            <pre><code>cd                     # Перейти в домашний каталог  
cd ~                   # Перейти в домашний каталог (альтернатива)  
cd /путь/к/каталогу    # Перейти в указанный каталог (абсолютный путь)  
cd каталог             # Перейти в подкаталог (относительный путь)  
cd ..                  # Подняться на уровень выше  
cd ../..               # Подняться на два уровня выше  
cd -                   # Вернуться в предыдущий каталог</code></pre>

            <p><strong>Полезные сокращения:</strong></p>
            <ul>
                <li><code>~</code> - домашний каталог пользователя (/home/username)</li>
                <li><code>.</code> - текущий каталог</li>
                <li><code>..</code> - родительский каталог</li>
            </ul>
        </div>

        <div class="section">
            <h2>📁 Просмотр содержимого каталогов</h2>
            
            <h3>Базовые команды просмотра</h3>
            <pre><code>ls                     # Показать файлы и каталоги  
ls -l                  # Подробный список с правами и размерами  
ls -la                 # Показать все файлы (включая скрытые)  
ls -lh                 # Размеры в удобном формате (K, M, G)  
ls -t                  # Сортировка по времени изменения  
ls -S                  # Сортировка по размеру</code></pre>
        </div>

        <!-- Остальные секции будут размещены аналогично -->

        <div class="section">
            <h2>🗺️ Структура файловой системы Linux</h2>
            
            <h3>Основные каталоги в корне системы (/)</h3>
            <ul>
                <li><code>/home</code> - домашние каталоги пользователей</li>
                <li><code>/etc</code> - конфигурационные файлы системы</li>
                <li><code>/bin</code> - основные исполняемые программы</li>
                <li><code>/usr</code> - пользовательские программы и данные</li>
                <li><code>/var</code> - изменяемые данные (логи, кэш)</li>
                <li><code>/tmp</code> - временные файлы</li>
                <li><code>/dev</code> - файлы устройств</li>
                <li><code>/media</code> - точки монтирования съемных устройств</li>
            </ul>
        </div>

        <div class="section">
            <h2>⚠️ Важные замечания</h2>
            <ol>
                <li><strong>Чувствительность к регистру</strong>: Linux различает файлы <code>Файл.txt</code> и <code>файл.txt</code></li>
                <li><strong>Безопасность</strong>: Будьте осторожны с командами <code>rm -rf</code> - они могут безвозвратно удалить данные</li>
                <li><strong>Автодополнение</strong>: Активно используйте Tab для ускорения набора и избежания ошибок</li>
                <li><strong>Справка</strong>: Используйте <code>man команда</code> или <code>команда --help</code> для получения подробной справки</li>
            </ol>
        </div>
    </div>
</body>
</html>

