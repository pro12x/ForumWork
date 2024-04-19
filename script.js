class MonCanvas extends HTMLElement {
 constructor() {
    super();
    // Initialisation de l'élément
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.innerHTML = `
      <style>
        :host {
          display: block;
          width: 300px;
          height: 150px;
          border: 1px solid black;
        }
        #toile {
          width: 100%;
          height: 100%;
          background-color: white;
        }
      </style>
      <div id="toile"></div>
    `;
    this.toile = this.shadowRoot.querySelector('#toile');
    this.svg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
    this.svg.setAttribute('width', '100%');
    this.svg.setAttribute('height', '100%');
    this.toile.appendChild(this.svg);
 }

 getContext(contextType) {
    if (contextType === '2d') {
      return {
        fillStyle: 'black',
        strokeStyle: 'black',
        lineWidth: 1,
        fillRect: (x, y, width, height) => {
          const rect = document.createElementNS('http://www.w3.org/2000/svg', 'rect');
          rect.setAttribute('x', x);
          rect.setAttribute('y', y);
          rect.setAttribute('width', width);
          rect.setAttribute('height', height);
          rect.setAttribute('fill', this.fillStyle);
          this.svg.appendChild(rect);
        },
        drawCircle: (x, y, radius) => {
          const circle = document.createElementNS('http://www.w3.org/2000/svg', 'circle');
          circle.setAttribute('cx', x);
          circle.setAttribute('cy', y);
          circle.setAttribute('r', radius);
          circle.setAttribute('fill', this.fillStyle);
          this.svg.appendChild(circle);
        },
        // Ajoutez d'autres méthodes de dessin ici
      };
    }
    throw new Error('Seul le contexte 2D est supporté pour le moment.');
 }
}

customElements.define('mon-canvas', MonCanvas);
