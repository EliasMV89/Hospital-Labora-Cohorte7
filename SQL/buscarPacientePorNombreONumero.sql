 -- 2. Buscar pacientes por nombre, número de seguro social u otra información relevante
SELECT * FROM Pacientes WHERE Nombre LIKE '%nombre%' OR NumeroSeguroSocial = 'numero';
