const $ = document.querySelector.bind(document);
const $$ = document.querySelectorAll.bind(document);

document.addEventListener('DOMContentLoaded', () => {
  const imgInput = $('#imgInput');
  const output = $('#typeOption');
  const btn = $('button');
  const imgResult = $('#imgResult');
  const asciiPre = $('#ascii');

  asciiPre.classList.add('hide');

  btn.addEventListener('click', () => {
    const url = `${window.location.origin}/ascii.go?img=${
      imgInput.value
      }&output=${output[output.selectedIndex].value}`;

    if (output[output.selectedIndex].value !== 'ascii') {
      disableForm(true);

      imgResult.setAttribute('src', '');
      imgResult.setAttribute('src', url);

      imgResult.addEventListener('load', () => {
        disableForm(false);
      });
    } else {
      const loading = $('.js-loading');
      asciiPre.classList.remove('hide');
      imgResult.classList.add('hide');

      loading.classList.add('hide');
      getData(url).then(result => {
        asciiPre.innerText = result.ascii;
        loading.classList.remove('hide');
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
});
