window.onload = function(){
    document.querySelectorAll(".link").forEach(item => {
        item.addEventListener('transitionend', function(e) {
          e.stopPropagation();
        })
    });
}

function navigate(u) {
    var container = document.getElementsByClassName("container")[0];
    container.onanimationend = () => {
        window.location.href = u;
    }
    container.classList.add("fade");
}

