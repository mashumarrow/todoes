package graph

import (
    "context"
    "github.com/mashumarrow/todoes/models"
    "gorm.io/gorm"
)

type Resolver struct {
    DB *gorm.DB
}

func (r *Resolver) Query_subjects(ctx context.Context) ([]*models.Subject, error) {
    var subjects []*models.Subject
    if err := r.DB.Find(&subjects).Error; err != nil {
        return nil, err
    }
    return subjects, nil
}

func (r *Resolver) Query_schedules(ctx context.Context, dayOfWeek string) ([]*models.Schedule, error) {
    var schedules []*models.Schedule
    if err := r.DB.Where("day_of_week = ?", dayOfWeek).Preload("Subject").Preload("Classroom").Find(&schedules).Error; err != nil {
        return nil, err
    }
    return schedules, nil
}

func (r *Resolver) Mutation_createSubject(ctx context.Context, name string) (*models.Subject, error) {
    subject := &models.Subject{Name: name}
    if err := r.DB.Create(subject).Error; err != nil {
        return nil, err
    }
    return subject, nil
}

func (r *Resolver) Mutation_createClassroom(ctx context.Context, name string) (*models.Classroom, error) {
    classroom := &models.Classroom{Name: name}
    if err := r.DB.Create(classroom).Error; err != nil {
        return nil, err
    }
    return classroom, nil
}

func (r *Resolver) Mutation_createSchedule(ctx context.Context, subjectID uint, classroomID uint, dayOfWeek string, period int) (*models.Schedule, error) {
    schedule := &models.Schedule{
        SubjectID:  subjectID,
        ClassroomID: classroomID,
        DayOfWeek:  dayOfWeek,
        Period:     period,
    }
    if err := r.DB.Create(schedule).Error; err != nil {
        return nil, err
    }
    return schedule, nil
}

func (r *Resolver) Mutation_createTodo(ctx context.Context, subjectID uint, title string) (*models.Todo, error) {
    todo := &models.Todo{SubjectID: subjectID, Title: title}
    if err := r.DB.Create(todo).Error; err != nil {
        return nil, err
    }
    return todo, nil
}

func (r *Resolver) Mutation_toggleTodoComplete(ctx context.Context, todoID uint) (*models.Todo, error) {
    var todo models.Todo
    if err := r.DB.First(&todo, todoID).Error; err != nil {
        return nil, err
    }
    todo.Completed = !todo.Completed
    if err := r.DB.Save(&todo).Error; err != nil {
        return nil, err
    }
    return &todo, nil
}
