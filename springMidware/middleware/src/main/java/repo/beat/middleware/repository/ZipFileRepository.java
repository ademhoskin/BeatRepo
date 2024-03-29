package repo.beat.middleware.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import repo.beat.middleware.model.BeatZip;

public interface ZipFileRepository extends JpaRepository<BeatZip, Long>{
    
}
