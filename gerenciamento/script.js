function updateProgress(index, value) {
    // Atualiza a largura da barra de progresso
    const progressBar = document.querySelectorAll('.progress')[index];
    progressBar.style.width = value + '%';

    // Atualiza o texto da porcentagem
    const percentageText = document.querySelectorAll('.percentage')[index];
    percentageText.textContent = value + '%';
}