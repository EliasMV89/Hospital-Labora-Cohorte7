-- 4. Mostrar las citas próximas para un médico específico
SELECT * FROM Citas WHERE IDMedico = idMedico AND FechaHora > NOW() ORDER BY FechaHora;
