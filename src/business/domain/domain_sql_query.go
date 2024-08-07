package domain

// Area
const (
	_sqlGetAreaByID  = `SELECT id, name, latitude, longitude, code_area, notes, created_at, updated_at FROM area WHERE status=1 AND id=? LIMIT 1`
	_sqlCreateArea   = `INSERT INTO area (name, latitude, longitude, code_area, notes) VALUES (?,?,?,?,?)`
	_sqlGetArea      = `SELECT id, name, latitude, longitude, code_area, notes, created_at, updated_at FROM area WHERE status=1`
	_sqlGetAreaCount = `SELECT COUNT(*) FROM area WHERE status=1`
	_sqlUpdateArea   = `UPDATE area SET name = ?, code_area = ?, notes = ?, latitude = ?, longitude = ? WHERE id = ?`
	_sqlDeleteArea   = `UPDATE area SET status=0 WHERE id = ?`
)
