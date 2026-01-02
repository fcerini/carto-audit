package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func iniciarProceso() {
	// Truncate a la tabla auditoria.procesar
	_, err := DB.Exec("TRUNCATE TABLE auditoria.procesar")
	if err != nil {
		log.Fatalf("Failed to truncate table: %v", err)
	}
	log.Println("Truncated table auditoria.procesar successfully.")

	// Ejecutar gloAPP.SQLCartas para obtener las cartas a procesar
	rows, err := DB.Query(gloAPP.SQLCartas)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	idpk := 0
	for rows.Next() {
		idpk++
		// Procesar cada fila, insertar en auditoria.procesar
		suceso := sql.NullString{}
		provincia := sql.NullString{}
		partido := sql.NullString{}
		localidad := sql.NullString{}
		calle := sql.NullString{}
		altura := sql.NullString{}
		callesuperior := sql.NullString{}
		calleinferior := sql.NullString{}
		err := rows.Scan(&suceso, &provincia, &partido, &localidad, &calle, &altura, &callesuperior, &calleinferior)
		if err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}

		_, err = DB.Exec(`INSERT INTO auditoria.procesar 
			(idpk, suceso, provincia, partido, localidad, calle, altura, callesuperior, calleinferior) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
			idpk, getString(suceso), getString(provincia), getString(partido), getString(localidad), getString(calle), getString(altura), getString(callesuperior), getString(calleinferior))
		if err != nil {
			log.Fatalf("Failed to insert into auditoria.procesar: %v", err)
		}

		if idpk%100 == 0 {
			log.Printf("Processed %d records...", idpk)
		}

	}
	log.Printf("Proceso completado. Total de registros procesados: %d", idpk)

}
