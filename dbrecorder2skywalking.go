package cls_skywalking_client_go

import (
	"github.com/Masterminds/squirrel"
	"github.com/Masterminds/structable"
	"fmt"
	"log"
)

type RecordProxy struct {
	Recorder structable.Recorder
}

// “构造基类”
func NewRecorderProxy(recorder structable.Recorder) *RecordProxy {
	return &RecordProxy{
		Recorder: recorder,
	}
}

func (f RecordProxy) DB() squirrel.DBProxyBeginner {
	return f.Recorder.DB()
}

func (f RecordProxy) getRecorder() structable.Recorder {
	return f.Recorder
}

func (f *RecordProxy) Insert() error {
	queryStr := fmt.Sprintf("Insert table %s, whereIds: %v", f.Recorder.TableName(), f.Recorder.WhereIds())
	reqSpan, spanErr := StartSpantoSkyWalkingForDb(queryStr, GetDsn(f.DB()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb insert error: %v \n", spanErr)
	}

	err := f.getRecorder().Insert()

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, false, err)
	} else {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, true, nil)
	}
	return err
}

func (f *RecordProxy) Delete() error {
	queryStr := fmt.Sprintf("Delete table %s, whereIds: %v", f.Recorder.TableName(), f.Recorder.WhereIds())
	reqSpan, spanErr := StartSpantoSkyWalkingForDb(queryStr, GetDsn(f.DB()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb Delete error: %v \n", spanErr)
	}

	err := f.getRecorder().Delete()

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, false, err)
	} else {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, true, nil)
	}
	return err
}

func (f *RecordProxy) Update() error {
	queryStr := fmt.Sprintf("Update table %s, whereIds: %v", f.Recorder.TableName(), f.Recorder.WhereIds())
	reqSpan, spanErr := StartSpantoSkyWalkingForDb(queryStr, GetDsn(f.DB()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb Update error: %v \n", spanErr)
	}
	err := f.getRecorder().Update()

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, false, err)
	} else {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, true, nil)
	}
	return err
}

func (f *RecordProxy) Load() error {
	queryStr := fmt.Sprintf("Load table %s, whereIds: %v", f.Recorder.TableName(), f.Recorder.WhereIds())
	reqSpan, spanErr := StartSpantoSkyWalkingForDb(queryStr, GetDsn(f.DB()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb Load error: %v \n", spanErr)
	}

	err := f.getRecorder().Load()

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, false, err)
	} else {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, true, nil)
	}
	return err
}

func (f *RecordProxy) LoadWhere(arg1 interface{}, arg2 ...interface{}) error {
	queryStr := fmt.Sprintf("LoadWhere table %s, whereIds: %v", f.Recorder.TableName(), f.Recorder.WhereIds())
	reqSpan, spanErr := StartSpantoSkyWalkingForDb(queryStr, GetDsn(f.DB()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb LoadWhere error: %v \n", spanErr)
	}

	err := f.getRecorder().LoadWhere(arg1, arg2...)

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, false, err)
	} else {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, true, nil)
	}
	return err
}

func (f *RecordProxy) ExistsWhere(arg1 interface{}, arg2 ...interface{}) (bool, error) {
	queryStr := fmt.Sprintf("ExistsWhere table %s, whereIds: %v", f.Recorder.TableName(), f.Recorder.WhereIds())
	reqSpan, spanErr := StartSpantoSkyWalkingForDb(queryStr, GetDsn(f.DB()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb ExistsWhere error: %v \n", spanErr)
	}

	boolValue, err := f.getRecorder().ExistsWhere(arg1, arg2...)

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, false, err)
	} else {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, true, nil)
	}
	return boolValue, err
}

func (f *RecordProxy) Columns(boolean bool) []string {
	return f.getRecorder().Columns(boolean)
}

func (f *RecordProxy) FieldReferences(boolean bool) []interface{} {
	return f.getRecorder().FieldReferences(boolean)
}

func (f *RecordProxy) Exists() (bool, error) {
	queryStr := fmt.Sprintf("Exists table %s, whereIds: %v", f.Recorder.TableName(), f.Recorder.WhereIds())
	reqSpan, spanErr := StartSpantoSkyWalkingForDb(queryStr, GetDsn(f.DB()))
	if spanErr != nil {
		log.Printf("StartSpantoSkyWalkingForDb Exists error: %v \n", spanErr)
	}

	boolValue, err := f.getRecorder().Exists()

	if err != nil {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, false, err)
	} else {
		EndSpantoSkywalkingForDb(reqSpan, queryStr, true, nil)
	}
	return boolValue, err
}

