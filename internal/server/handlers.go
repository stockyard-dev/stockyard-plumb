package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-plumb/internal/store")
func(s *Server)handleListDBs(w http.ResponseWriter,r *http.Request){list,_:=s.db.ListDatabases();if list==nil{list=[]store.Database{}};writeJSON(w,200,list)}
func(s *Server)handleCreateDB(w http.ResponseWriter,r *http.Request){var d store.Database;json.NewDecoder(r.Body).Decode(&d);if d.Name==""{writeError(w,400,"name required");return};s.db.CreateDatabase(&d);writeJSON(w,201,d)}
func(s *Server)handleDeleteDB(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.Delete(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleListMigrations(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);list,_:=s.db.ListMigrations(id);if list==nil{list=[]store.Migration{}};writeJSON(w,200,list)}
func(s *Server)handleRecordMigration(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var m store.Migration;json.NewDecoder(r.Body).Decode(&m);m.DatabaseID=id;if m.Version==""||m.Status==""{writeError(w,400,"version and status required");return};s.db.RecordMigration(&m);writeJSON(w,201,m)}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}
