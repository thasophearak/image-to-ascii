package ui

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

func Index(w http.ResponseWriter) {
	template := `
	<!DOCTYPE html>
	<html lang="en">
		 <head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<meta http-equiv="X-UA-Compatible" content="ie=edge">
				<title>Image to ASCII</title>
				<meta name="twitter:image" content="https://image-to-ascii.now.sh/assets/image-2-ascii.png">
				<meta property="og:image" content="https://image-to-ascii.now.sh/assets/image-2-ascii.png">
				<style type="text/css">
					*,:after,:before {
						box-sizing: border-box
					}
					body {
						font-family: -apple-system,BlinkMacSystemFont,San Francisco,Helvetica Neue,Helvetica,Ubuntu,Google Sans,Roboto,Noto,Segoe UI,Arial,sans-serif;
						margin: 0;
						color: #111
					}
					h1 {
						font-size: 32px;
						line-height: 42px;
						text-align: center;
						margin: 32px 0
					}
					a {
						cursor: pointer;
						color: #0076ff;
						text-decoration: none;
						transition: all .2s ease;
						border-bottom: 1px solid #fff
					}
					a:hover {
						border-bottom: 1px solid #0076ff
					}
					button {
						-webkit-appearance: none;
						width: 100%;
						border: 1px solid rgba(0,0,0,.12);
						padding: 12px;
						border-radius: 7px;
						font-size: inherit;
						background: #111;
						color: #fff;
						cursor: pointer
					}
					button:disabled {
						color: #ccc;
						background: rgba(0,0,0,.08);
						border: 1px solid transparent;
						cursor: progress
					}
					svg {
						width: 24px;
						height: 24px;
						stroke: currentColor;
						stroke-width: 2;
						stroke-linecap: round;
						stroke-linejoin: round;
						fill: none
					}
					pre {
						white-space: nowrap;
						overflow: scroll;
						font-family: monospace,monospace;
						padding: 12px;
						margin: 0;
						border: 1px solid rgba(0,0,0,.12);
						border-radius: 7px
					}
					.hide {
						display: none!important
					}
					.container {
						display: flex;
						flex-flow: column;
						width: 100%;
						margin: 0 auto;
						padding: 0 24px
					}
					.fields {
						width: 100%;
						margin-bottom: 24px
					}
					.fields label {
						margin-bottom: 12px
					}
					.fields input,.fields select {
						-webkit-appearance: none;
						outline: 0;
						width: 100%;
						max-width: 100%;
						background: #fff;
						height: 40px;
						border: 1px solid rgba(0,0,0,.12);
						border-radius: 7px;
						padding: 0 12px;
						font-size: inherit;
						transition: all .15s ease-in
					}
					.fields input:active,.fields input:focus,.fields select:active,.fields select:focus {
						outline: none;
						box-shadow: 0 1px 4px 0 rgba(0,0,0,.08)
					}
					.fields input:disabled,.fields select:disabled {
						cursor: progress;
						color: #ccc;
						background: rgba(0,0,0,.08);
						border: 1px solid transparent
					}
					.fields input:disabled+svg,.fields select:disabled+svg {
						color: #ccc
					}
					.fields .with-icon {
						position: relative;
						width: 100%
					}
					.fields .with-icon svg {
						position: absolute;
						right: 8px;
						top: 8px
					}
					.fields .note p {
						font-size: 80%;
						font-style: italic;
						margin: 0
					}
					.field {
						display: flex;
						align-items: flex-start;
						justify-content: start;
						flex-flow: column
					}
					.field:not(:last-child) {
						margin-bottom: 24px
					}
					.result {
						width: 100%
					}
					.result img {
						box-shadow: 0 1px 5px 0 rgba(0,0,0,.08);
						border-radius: 7px;
						width: 100%;
						height: auto
					}
					footer {
						text-align: center;
						padding: 48px 12px 24px
					}
					small {
						font-size: 80%;
						opacity: .6
					}
					@media screen and (min-width: 768px) {
						button {
							width:240px
						}
						.container {
							flex-flow: row
						}
						.fields {
							width: 40%;
							min-width: 410px;
							padding-right: 24px
						}
						.fields label {
							margin-bottom: 0;
							margin-right: 12px
						}
						.fields input,.fields select {
							width: 240px
						}
						.fields .with-icon {
							width: auto
						}
						.result {
							width: 60%;
							max-width: 800px
						}
						.field {
							align-items: center;
							justify-content: flex-end;
							flex-flow: row
						}
					}
					.ph-item {
						position: relative;
						display: -webkit-box;
						display: -ms-flexbox;
						display: flex;
						-ms-flex-wrap: wrap;
						flex-wrap: wrap;
						padding: 30px 15px 15px;
						overflow: hidden;
						margin-bottom: 30px;
						background-color: #fff;
						border: 1px solid #e6e6e6;
						border-radius: 7px
					}
					.ph-item,.ph-item *,.ph-item :after,.ph-item :before {
						-webkit-box-sizing: border-box;
						box-sizing: border-box
					}
					.ph-item:before {
						content: " ";
						position: absolute;
						top: 0;
						right: 0;
						bottom: 0;
						left: 50%;
						z-index: 1;
						width: 500%;
						margin-left: -250%;
						-webkit-animation: phAnimation .8s linear infinite;
						animation: phAnimation .8s linear infinite;
						background: -webkit-gradient(linear,left top,right top,color-stop(46%,hsla(0,0%,100%,0)),color-stop(50%,hsla(0,0%,100%,.35)),color-stop(54%,hsla(0,0%,100%,0))) 50% 50%;
						background: linear-gradient(90deg,hsla(0,0%,100%,0) 46%,hsla(0,0%,100%,.35) 50%,hsla(0,0%,100%,0) 54%) 50% 50%
					}
					.ph-item>* {
						-webkit-box-flex: 1;
						-ms-flex: 1 1 auto;
						flex: 1 1 auto;
						-webkit-box-orient: vertical;
						-webkit-box-direction: normal;
						-ms-flex-flow: column;
						flex-flow: column;
						padding-right: 15px;
						padding-left: 15px
					}
					.ph-item>*,.ph-row {
						display: -webkit-box;
						display: -ms-flexbox;
						display: flex
					}
					.ph-row {
						-ms-flex-wrap: wrap;
						flex-wrap: wrap;
						margin-bottom: 7.5px
					}
					.ph-row div {
						height: 10px;
						margin-bottom: 7.5px;
						background-color: #ced4da
					}
					.ph-row .big,.ph-row.big div {
						height: 20px;
						margin-bottom: 15px
					}
					.ph-row .empty {
						background-color: hsla(0,0%,100%,0)
					}
					.ph-col-12 {
						-webkit-box-flex: 0;
						-ms-flex: 0 0 100%;
						flex: 0 0 100%
					}
					.ph-picture {
						border-radius: 7px;
						width: 100%;
						height: 320px;
						background-color: #ccc;
						margin-bottom: 15px
					}
					@-moz-keyframes phAnimation {
						0% {
							-webkit-transform: translate3d(-30%,0,0);
							transform: translate3d(-30%,0,0)
						}
						to {
							-webkit-transform: translate3d(30%,0,0);
							transform: translate3d(30%,0,0)
						}
					}
					@-webkit-keyframes phAnimation {
						0% {
							-webkit-transform: translate3d(-30%,0,0);
							transform: translate3d(-30%,0,0)
						}
						to {
							-webkit-transform: translate3d(30%,0,0);
							transform: translate3d(30%,0,0)
						}
					}
					@-o-keyframes phAnimation {
						0% {
							-webkit-transform: translate3d(-30%,0,0);
							transform: translate3d(-30%,0,0)
						}
						to {
							-webkit-transform: translate3d(30%,0,0);
							transform: translate3d(30%,0,0)
						}
					}
					@keyframes phAnimation {
						0% {
							-webkit-transform: translate3d(-30%,0,0);
							transform: translate3d(-30%,0,0)
						}
						to {
							-webkit-transform: translate3d(30%,0,0);
							transform: translate3d(30%,0,0)
						}
					}			
				</style>
		 </head>
		 <body>
				<h1>Image to ASCII</h1>
				<div class="container">
					 <div class="fields">
							<div class="field"><label for="imgInput">Image</label><input class="js-input" type="text" id="imgInput" value="https://image-to-ascii.now.sh/assets/gopher.png"></div>
							<div class="field">
								 <label for="typeOption">Output</label>
								 <div class="with-icon">
										<select class="js-input" name="typeOption" id="typeOption">
											 <option value="ascii">ASCII</option>
											 <option value="png">PNG</option>
											 <option value="jpg" selected>JPG</option>
										</select>
										<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24">
											 <path d="M6 9l6 6 6-6"/>
										</svg>
								 </div>
							</div>
							<div class="field"><button class="js-button">Generate</button></div>
							<div class="field">
								 <div class="note">
										<p class="hide" id="js-copied">Copied to your clipboard!</p>
								 </div>
							</div>
					 </div>
					 <div class="result">
							<pre id="ascii"></pre>
							<img src="/assets/gopher-ascii.jpeg" id="imgResult">
							<div class="hide ph-item js-loading">
								 <div class="ph-col-12">
										<div class="ph-picture"></div>
								 </div>
							</div>
					 </div>
				</div>
				<footer>
					 <h2>What is this?</h2>
					 <p>This is a service that generate ascii image.</p>
					 <p>Find out how it work <a href="https://github.com/sophearak/image-to-ascii">GitHub</a></p>
					 <small>Proudly hosted on <a href="https://zeit.co/now">▲ZEIT Now</a></small>
				</footer>

				<script type="text/javascript">
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
							const url = window.location.origin + '/?img=' + imgInput.value + '&output=' + output[output.selectedIndex].value;
					
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
					
						function getData(url = '') {
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
				</script>
		 </body>
	</html>
	`
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	m.Add("text/html", &html.Minifier{
		KeepDocumentTags: true,
	})

	mt, err := m.String("text/html", template)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", strconv.Itoa(len(mt)))
	w.Write([]byte(mt))
}
