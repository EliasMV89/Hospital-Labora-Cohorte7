-- 6. Generar un informe de pacientes atendidos por un médico en un período de tiempo específico
SELECT * FROM Citas WHERE IDMedico = idMedico AND FechaHora BETWEEN '2024-06-01' AND '2024-06-30';
