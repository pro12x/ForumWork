<script>
 const monCanvas = document.querySelector('mon-canvas');
 const ctx = monCanvas.getContext('2d');

 // Dessiner un rectangle
 ctx.fillStyle = 'red';
 ctx.fillRect(10, 10, 50, 50);

 // Dessiner un cercle
 ctx.fillStyle = 'blue';
 ctx.drawCircle(100, 100, 20);
</script>
