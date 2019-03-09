const $ = document.querySelector.bind(document);
const $$ = document.querySelectorAll.bind(document);

document.addEventListener('DOMContentLoaded', () => {
  const imgInput = $('#imgInput');
  const output = $('#typeOption');
  const btn = $('button');
  const imgResult = $('#imgResult');
  const asciiPre = $('#ascii');
  const copied = $('#js-copied')

  asciiPre.classList.add('hide');

  btn.addEventListener('click', () => {
    const url = `${window.location.origin}/ascii.go?img=${
      imgInput.value
      }&output=${output[output.selectedIndex].value}`;

    if (output[output.selectedIndex].value !== 'ascii') {
      disableForm(true);

      imgResult.setAttribute('src', '');
      imgResult.setAttribute('src', url);
      copied.classList.add('hide');
      asciiPre.classList.add('hide');
      imgResult.addEventListener('load', () => {
        disableForm(false);
      });
    } else {
      const loading = $('.js-loading');
      imgResult.classList.add('hide');

      loading.classList.remove('hide');
      copied.classList.add('hide');
      getData(url).then(result => {
        asciiPre.classList.remove('hide');
        asciiPre.innerText = result.ascii;
        loading.classList.add('hide');
        copyToClipboard(result.ascii);
        copied.classList.remove('hide');
      });
    }
  });

  function disableForm(disable) {
    const inputs = $$('.js-input');
    const btn = $('.js-button');
    const imgResult = $('#imgResult');
    const loading = $('.js-loading');

    if (disable) {
      btn.setAttribute('disabled', true);
      imgResult.classList.add('hide');
      loading.classList.remove('hide');
      return inputs.forEach(input => {
        input.setAttribute('disabled', 'true');
      });
    }

    loading.classList.add('hide');
    btn.removeAttribute('disabled');
    imgResult.classList.remove('hide');
    inputs.forEach(input => {
      input.removeAttribute('disabled');
    });
  }

  function getData(url = ``) {
    return fetch(url, {
      method: 'GET',
      cache: 'no-cache',
      headers: {
        'Content-Type': 'application/json',
      }
    })
      .then(response => response.json());
  }

  const copyToClipboard = str => {
    const el = document.createElement('textarea');
    el.style.whiteSpace = 'pre'
    el.setAttribute('readonly', '');
    el.style.position = 'absolute';
    el.style.left = '-9999px';
    el.value = str;
    document.body.appendChild(el);
    el.select();
    document.execCommand('copy');
    document.body.removeChild(el);
  };
});
