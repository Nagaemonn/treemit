document.addEventListener('DOMContentLoaded', function() {
    const li = document.getElementById('dark-mode-toggle');
    if (!li) return;
    const label = li.querySelector('span');
    function updateLabel() {
      const isDark = document.documentElement.dataset.scheme === 'dark';
      label.textContent = isDark ? 'Light Mode' : 'Dark Mode';
    }
    // 既存のクリックイベントの後にラベルも更新
    li.addEventListener('click', function() {
      setTimeout(updateLabel, 10); // テーマ切り替え後に実行
    });
    // 初期表示
    updateLabel();
  });