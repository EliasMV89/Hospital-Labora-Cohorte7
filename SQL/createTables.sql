CREATE TABLE Pacientes (
                           ID INT PRIMARY KEY,
                           Nombre VARCHAR(100),
                           NumeroSeguroSocial VARCHAR(50),
                           GravedadSalud INT
);

CREATE TABLE Medicos (
                         ID INT PRIMARY KEY,
                         Nombre VARCHAR(100)
);

CREATE TABLE Citas (
                       ID INT PRIMARY KEY,
                       IDPaciente INT,
                       IDMedico INT,
                       FechaHora DATETIME,
                       FOREIGN KEY (IDPaciente) REFERENCES Pacientes(ID),
                       FOREIGN KEY (IDMedico) REFERENCES Medicos(ID)
);

CREATE TABLE HistorialMedico (
                                 ID INT PRIMARY KEY,
                                 IDPaciente INT,
                                 Diagnostico VARCHAR(255),
                                 Fecha DATETIME,
                                 FOREIGN KEY (IDPaciente) REFERENCES Pacientes(ID)
);
